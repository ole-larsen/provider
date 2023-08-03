package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/olelarssen/provider/models"
	"github.com/olelarssen/provider/service/settings"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"
)

const oauthYandexUrlAPI = "https://oauth.yandex.com/authorize?response_type=token&client_id="

var yandexOauthConfig = &oauth2.Config{
	RedirectURL:  settings.Settings.Auth.Yandex.Callback,
	ClientID:     settings.Settings.Auth.Yandex.ClientID,
	ClientSecret: settings.Settings.Auth.Yandex.ClientSecret,
	Endpoint:     yandex.Endpoint,
}

func getUserDataFromYandex(code string) (*models.UserInfo, error) {
	// Use code to get token and get user info from Google.

	token, err := yandexOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	fmt.Println(token)
	response, err := http.Get(oauthYandexUrlAPI + settings.Settings.Auth.Yandex.ClientID)
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