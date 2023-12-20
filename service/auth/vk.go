package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func getUserDataFromVk(code string, state string) (*models.VkUserInfo, error) {
	// Use code to get token and get user info from Google.
	fmt.Println(code, state)
	/*
		https://oauth.vk.com/access_token?client_id=1&client_secret=H2Pk8htyFD8024mZaPHm&redirect_uri=http://mysite.ru&code=7a6fa4dff77a228eeda56603b8f53806c883f011c40b72630bb50df056f6479e52a
	*/
	oauthVkUrlAPI := vkOauthConfig.Endpoint.TokenURL + "?client_id=" + vkOauthConfig.ClientID + "&client_secret=" + "&redirect_url=" + vkOauthConfig.RedirectURL + "&code=" + code
	client := http.Client{}
	req, err := http.NewRequest("GET", oauthVkUrlAPI, nil)
	if err != nil {
		//Handle Error
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}

	response, err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed getting access token: %s", err.Error())
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	var vkUserInfo models.VkUserInfo

	err = json.Unmarshal([]byte(contents), &vkUserInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println(vkUserInfo)

	return &vkUserInfo, nil
}

func (s *Server) VkLogin(w http.ResponseWriter, p runtime.Producer) string {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	authURL := vkOauthConfig.AuthCodeURL(oauthState) + "&scope=friends&response_type=code&v=5.131"

	return authURL
}

func (s *Server) VkCallback(code string, state string) (*models.VkUserInfo, error) {
	return getUserDataFromVk(code, state)
}
