package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/olelarssen/provider/models"
	"github.com/olelarssen/provider/service/settings"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  settings.Settings.Auth.Google.Callback,
	ClientID:     settings.Settings.Auth.Google.ClientID,
	ClientSecret: settings.Settings.Auth.Google.ClientSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string) (*models.UserInfo, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	userInfo := models.UserInfo{}

	err = json.Unmarshal([]byte(contents), &userInfo)
	if err != nil {
		return nil, err
	}
	userInfo.Token = &models.Token{
		AccessToken:  &token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
	return &userInfo, nil
}

func (s *Server) GoogleLogin(w http.ResponseWriter, p runtime.Producer) string {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	return googleOauthConfig.AuthCodeURL(oauthState)
}

func (s *Server) GoogleCallback(code string) (*models.UserInfo, error) {
	return getUserDataFromGoogle(code)
}
