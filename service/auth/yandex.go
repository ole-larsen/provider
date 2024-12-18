package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/olelarssen/provider/models"
	"github.com/olelarssen/provider/service/settings"
	"golang.org/x/oauth2"
)

const oauthYandexUrlAPI = "https://login.yandex.ru/info?format=json"

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://oauth.yandex.ru/authorize",
	TokenURL: "https://oauth.yandex.ru/token",
}

var yandexOauthConfig = &oauth2.Config{
	RedirectURL:  settings.Settings.Auth.Yandex.Callback,
	ClientID:     settings.Settings.Auth.Yandex.ClientID,
	ClientSecret: settings.Settings.Auth.Yandex.ClientSecret,
	Endpoint:     Endpoint, //yandex.Endpoint,
	Scopes:       []string{"login:birthday", "login:email", "login:info", "login:avatar"},
}

func getUserDataFromYandex(code string) (*models.UserInfo, error) {
	// Use code to get token and get user info from Google.

	token, err := yandexOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", oauthYandexUrlAPI, nil)
	if err != nil {
		//Handle Error
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"OAuth " + token.AccessToken},
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
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
		RefreshToken: &token.RefreshToken,
	}
	return &userInfo, nil
}

func (s *Server) YandexLogin(w http.ResponseWriter, p runtime.Producer) string {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	return yandexOauthConfig.AuthCodeURL(oauthState)
}

func (s *Server) YandexCallback(code string) (*models.UserInfo, error) {
	return getUserDataFromYandex(code)
}
