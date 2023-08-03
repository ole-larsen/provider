package auth

import (
	"net/http"

	"github.com/go-oauth2/oauth2/generates"
	"github.com/go-oauth2/oauth2/manage"
	oauthModels "github.com/go-oauth2/oauth2/models"
	"github.com/go-oauth2/oauth2/server"
	"github.com/go-oauth2/oauth2/store"
	"github.com/go-session/session"
	"github.com/google/uuid"
	"github.com/olelarssen/provider/models"
	log "github.com/olelarssen/provider/service/logger"
	"gopkg.in/oauth2.v3/errors"
)

type Server struct {
	manager *manage.Manager
	store   *store.ClientStore
	Service *server.Server
	logger  log.Logger
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}
	sessionStore.Set("LoggedInUserID", r.Form.Get("client_id"))
	uid, _ := sessionStore.Get("LoggedInUserID")

	userID = uid.(string)
	sessionStore.Delete("LoggedInUserID")
	sessionStore.Save()
	return
}

func NewServer(logger log.Logger) *Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	// client memory store
	clientStore := store.NewClientStore()

	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logger.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		logger.Println("Response Error:", re.Error.Error())
	})

	return &Server{
		manager: manager,
		store:   clientStore,
		Service: srv,
		logger:  logger,
	}
}

func (s *Server) NewClient(domain string, clientID string) *models.Credentials {
	clientSecret := uuid.New().String()[:8]
	err := s.store.Set(clientID, &oauthModels.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: domain,
		UserID: clientID,
	})

	if err != nil {
		s.logger.Errorln(err)
	}
	return &models.Credentials{
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
		Domain:       &domain,
	}
}
