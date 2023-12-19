package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/olelarssen/provider/models"
	"github.com/olelarssen/provider/service/settings"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

// const oauthVKUrlAPI = "https://login.yandex.ru/info?format=json"

var vkOauthConfig = &oauth2.Config{
	RedirectURL:  settings.Settings.Auth.Vk.Callback,
	ClientID:     settings.Settings.Auth.Vk.ClientID,
	ClientSecret: settings.Settings.Auth.Vk.ClientSecret,
	Endpoint:     vk.Endpoint,
	Scopes:       []string{},
}

func getUserDataFromVk(code string) (*models.UserInfo, error) {
	// Use code to get token and get user info from Google.

	token, err := vkOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	fmt.Println(code, token)
	// client, err := vk.NewClientWithOptions(vk.WithToken(token.AccessToken))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//user := getCurrentUser(client)

	// client := http.Client{}
	// req, err := http.NewRequest("GET", oauthVKUrlAPI, nil)
	// if err != nil {
	// 	//Handle Error
	// 	return nil, err
	// }

	// req.Header = http.Header{
	// 	"Content-Type":  {"application/json"},
	// 	"Authorization": {"OAuth " + token.AccessToken},
	// }

	// response, err := client.Do(req)
	// defer response.Body.Close()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	// }

	// contents, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed read response: %s", err.Error())
	// }

	userInfo := models.UserInfo{}

	// err = json.Unmarshal([]byte(contents), &userInfo)
	// if err != nil {
	// 	return nil, err
	// }
	userInfo.Token = &models.Token{
		AccessToken:  &token.AccessToken,
		RefreshToken: &token.RefreshToken,
	}
	return &userInfo, nil
}

func (s *Server) VkLogin(w http.ResponseWriter, p runtime.Producer) string {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	fmt.Println(vkOauthConfig)
	fmt.Println(settings.Settings.Auth.Vk)
	authURL := vkOauthConfig.AuthCodeURL(oauthState) + "scope=12"

	return authURL
}

func (s *Server) VkCallback(code string) (*models.UserInfo, error) {
	return getUserDataFromVk(code)
}
