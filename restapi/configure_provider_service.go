// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-session/session"
	"github.com/olelarssen/provider/restapi/operations/instruments"
	"github.com/olelarssen/provider/service/auth"
	"github.com/olelarssen/provider/service/metrics"
	"github.com/olelarssen/provider/service/settings"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/olelarssen/provider/restapi/operations"
	"github.com/olelarssen/provider/restapi/operations/public"
	log "github.com/olelarssen/provider/service/logger"
)

//go:generate swagger generate metrics --target ../../provider-service --name ProviderService --spec ../schema/swagger.yml --principal interface{}

func configureFlags(api *operations.ProviderServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("\n" + header + ": \n"))
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Payload struct {
	User User `json:"user"`
}

func configureAPI(api *operations.ProviderServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	logger := log.NewLogger()

	api.Logger = logger.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	s := auth.NewServer(logger)

	if api.PublicGetPingHandler == nil {
		api.PublicGetPingHandler = public.GetPingHandlerFunc(func(params public.GetPingParams) middleware.Responder {
			return middleware.NotImplemented("operation public.GetPing has not yet been implemented")
		})
	}

	api.PublicGetGoogleLoginHandler = public.GetGoogleLoginHandlerFunc(func(params public.GetGoogleLoginParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "google.login", params.HTTPRequest)
			u := s.GoogleLogin(w, p)
			http.Redirect(w, params.HTTPRequest, u, http.StatusFound)
		})
	})

	api.PublicGetGoogleCallbackHandler = public.GetGoogleCallbackHandlerFunc(func(params public.GetGoogleCallbackParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "google.callback", params.HTTPRequest)

			// Read oauthState from Cookie
			oauthState, _ := params.HTTPRequest.Cookie("oauthstate")

			if params.HTTPRequest.FormValue("state") != oauthState.Value {
				err := fmt.Errorf("invalid oauth google state")
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			userInfo, err := s.GoogleCallback(params.HTTPRequest.FormValue("code"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			logger.Println(userInfo)
			if userInfo != nil && userInfo.Token != nil {
				verified := "false"
				if userInfo.VerifiedEmail {
					verified = "true"
				}
				tokenType := ""
				if userInfo.Token.TokenType != nil {
					tokenType = *userInfo.Token.TokenType
				}
				scope := ""
				if userInfo.Token.Scope != nil {
					scope = *userInfo.Token.Scope
				}
				refreshToken := ""
				if userInfo.Token.RefreshToken != nil {
					refreshToken = *userInfo.Token.RefreshToken
				}
				expiresIn := ""
				if userInfo.Token.ExpiresIn != nil {
					expiresIn = fmt.Sprintf("%f", *userInfo.Token.ExpiresIn)
				}
				accessToken := ""
				if userInfo.Token.AccessToken != nil {
					accessToken = *userInfo.Token.AccessToken
				}

				if accessToken != "" {
					queryParams := url.Values{
						"client_id":      {userInfo.ClientID},
						"email":          {userInfo.Email},
						"id":             {userInfo.ID},
						"name":           {userInfo.Name},
						"picture":        {userInfo.Picture},
						"access_token":   {accessToken},
						"expires_in":     {expiresIn},
						"refresh_token":  {refreshToken},
						"scope":          {scope},
						"token_type":     {tokenType},
						"verified_email": {verified},
					}
					url := settings.Settings.Auth.Google.Redirect + "?" + queryParams.Encode() //

					http.RedirectHandler(url, http.StatusPermanentRedirect)
					return
				}
			}

			out, err := json.Marshal(userInfo)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(out))
			e := json.NewEncoder(w)
			e.SetIndent("", "  ")
			e.Encode(userInfo)
		})
	})

	api.PublicGetYandexLoginHandler = public.GetYandexLoginHandlerFunc(func(params public.GetYandexLoginParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "yandex.login", params.HTTPRequest)
			u := s.YandexLogin(w, p)
			http.Redirect(w, params.HTTPRequest, u, http.StatusFound)
		})
	})

	api.PublicGetYandexCallbackHandler = public.GetYandexCallbackHandlerFunc(func(params public.GetYandexCallbackParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "yandex.callback", params.HTTPRequest)

			// Read oauthState from Cookie
			oauthState, _ := params.HTTPRequest.Cookie("oauthstate")

			if params.HTTPRequest.FormValue("state") != oauthState.Value {
				err := fmt.Errorf("invalid oauth yandex state")
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			userInfo, err := s.YandexCallback(params.HTTPRequest.FormValue("code"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			e := json.NewEncoder(w)
			e.SetIndent("", "  ")
			e.Encode(userInfo)
		})
	})

	api.PublicGetVkLoginHandler = public.GetVkLoginHandlerFunc(func(params public.GetVkLoginParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "vk.login", params.HTTPRequest)
			u := s.VkLogin(w, p)
			http.Redirect(w, params.HTTPRequest, u, http.StatusFound)
		})
	})

	api.PublicGetVkCallbackHandler = public.GetVkCallbackHandlerFunc(func(params public.GetVkCallbackParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "vk.callback", params.HTTPRequest)

			request := params.HTTPRequest

			// Read oauthState from Cookie
			oauthState, _ := request.Cookie("oauthstate")

			if params.State != oauthState.Value {
				err := fmt.Errorf("invalid vk state")
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			userInfo, err := s.VkCallback(params.Code)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			e := json.NewEncoder(w)
			e.SetIndent("", "  ")
			e.Encode(userInfo)
		})
	})

	api.PublicGetTelegramLoginHandler = public.GetTelegramLoginHandlerFunc(func(params public.GetTelegramLoginParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "telegram.login", params.HTTPRequest)

			u := s.TelegramLogin(w, p)

			http.Redirect(w, params.HTTPRequest, u, http.StatusFound)
		})
	})

	api.PublicGetTelegramCallbackHandler = public.GetTelegramCallbackHandlerFunc(func(params public.GetTelegramCallbackParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "telegram.callback", params.HTTPRequest)

			r := params.HTTPRequest

			hash := r.URL.Query().Get("hash")
			payloadB64 := r.URL.Query().Get("payload")
			payloadBytes, err := base64.StdEncoding.DecodeString(payloadB64)
			if err != nil {
				http.Error(w, "Invalid payload", http.StatusBadRequest)
				return
			}

			var payload Payload
			err = json.Unmarshal(payloadBytes, &payload)
			if err != nil {
				http.Error(w, "Invalid payload", http.StatusBadRequest)
				return
			}

			h := hmac.New(sha256.New, []byte(settings.Settings.Auth.Telegram.PK+settings.Settings.Auth.Telegram.PK))
			h.Write([]byte(payloadB64))
			checkHash := hex.EncodeToString(h.Sum(nil))

			if hash != checkHash {
				http.Error(w, "Invalid hash", http.StatusBadRequest)
				return
			}

			user := payload.User
			//userId := user.Id
			// firstName := user.FirstName
			// lastName := user.LastName
			// username := user.Username
			logger.Println(user)
			// Store user information in your database
			// ...}

			e := json.NewEncoder(w)
			e.SetIndent("", "  ")
			e.Encode("")
		})
	})

	api.PublicGetCredentialsHandler = public.GetCredentialsHandlerFunc(func(params public.GetCredentialsParams) middleware.Responder {
		_ = dumpRequest(os.Stdout, "credentials", params.HTTPRequest)
		domain := settings.Settings.Domain
		if params.Domain != nil {
			domain = *params.Domain
		}
		payload := s.NewClient(domain, params.ClientID)
		logger.Infoln("domain:", domain, params.ClientID, payload)
		return public.NewGetCredentialsOK().WithPayload(payload)
	})

	api.PublicGetAuthorizeHandler = public.GetAuthorizeHandlerFunc(func(params public.GetAuthorizeParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "authorize", params.HTTPRequest)
			store, err := session.Start(params.HTTPRequest.Context(), w, params.HTTPRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var form url.Values
			if v, ok := store.Get("ReturnUri"); ok {
				form = v.(url.Values)
			}
			logger.Infoln(params.ClientID)

			if &params.RedirectURI != nil {
				logger.Infoln("RedirectURI:", params.RedirectURI)
			}
			if &params.Scope != nil {
				logger.Infoln("Scope:", params.Scope)
			}
			if params.State != nil {
				logger.Infoln("State", params.State)
			}
			logger.Infoln("Form:", form)
			logger.Infoln("HEADER:", params.HTTPRequest.Header.Get("Cf-Access-Authenticated-User-Email"))
			params.HTTPRequest.Form = form

			store.Delete("ReturnUri")
			store.Save()

			err = s.Service.HandleAuthorizeRequest(w, params.HTTPRequest)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		})
	})

	api.PublicPostTokenHandler = public.PostTokenHandlerFunc(func(params public.PostTokenParams) middleware.Responder {
		// http://localhost:5555/api/v1/token?client_id=222222&client_secret=22222222&domain=http://localhost:9094&grant_type=client_credentials&scope=read
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "token", params.HTTPRequest) // Ignore the error
			logger.Infoln(params.ClientID)
			if params.Scope != nil {
				logger.Infoln("Scope:", *params.Scope)
			}
			if params.Domain != nil {
				logger.Infoln("Domain:", *params.Domain)
			}
			if &params.ClientSecret != nil {
				logger.Infoln("ClientSecret:", params.ClientSecret)
			}
			if params.GrantType != nil {
				logger.Infoln("GrantType:", *params.GrantType)
			}
			logger.Infoln("Form:", params.HTTPRequest.Form)
			err := s.Service.HandleTokenRequest(w, params.HTTPRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		})
	})

	api.PublicGetValidateHandler = public.GetValidateHandlerFunc(func(params public.GetValidateParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "validate", params.HTTPRequest)
			token, err := s.Service.ValidationBearerToken(params.HTTPRequest)
			if params.AccessToken != nil {
				logger.Infoln(*params.AccessToken)
			}
			if err != nil {
				logger.Errorln(err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
				return
			}

			data := map[string]interface{}{
				"expires_in":   int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
				"client_id":    token.GetClientID(),
				"user_id":      token.GetUserID(),
				"access_token": token.GetAccess(),
			}
			e := json.NewEncoder(w)
			e.SetIndent("", "  ")
			e.Encode(data)
		})
	})

	api.InstrumentsGetMetricsHandler = instruments.GetMetricsHandlerFunc(func(params instruments.GetMetricsParams) middleware.Responder {
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			_ = dumpRequest(os.Stdout, "metrics", params.HTTPRequest)
			promhttp.Handler().ServeHTTP(w, params.HTTPRequest)
		})
	})
	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS metrics starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as metrics is initialized but not run yet, this function will be called.
// If you need to modify a config, store metrics instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return metrics.SetupHandler(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
