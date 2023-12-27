package auth

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/olelarssen/provider/models"
	"github.com/olelarssen/provider/service/settings"
)

// const oauthVKUrlAPI = "https://login.yandex.ru/info?format=json"

func getUserDataFromTelegram(code string) (*models.VkUserInfo, error) {
	// Use code to get token and get user info from Google.
	/*
		https://oauth.vk.com/access_token?client_id=1&client_secret=H2Pk8htyFD8024mZaPHm&redirect_uri=http://mysite.ru&code=7a6fa4dff77a228eeda56603b8f53806c883f011c40b72630bb50df056f6479e52a
	*/
	// oauthTelegramUrlAPI := telegramOauthConfig.Endpoint.TokenURL + "?client_id=" + telegramOauthConfig.ClientID + "&client_secret=" + telegramOauthConfig.ClientSecret + "&redirect_url=" + vkOauthConfig.RedirectURL + "&code=" + code
	// fmt.Println(oauthTelegramUrlAPI)

	// client := http.Client{}
	// req, err := http.NewRequest("GET", oauthTelegramUrlAPI, nil)
	// if err != nil {
	// 	//Handle Error
	// 	return nil, err
	// }

	// response, err := client.Do(req)
	// defer response.Body.Close()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed getting access token: %s", err.Error())
	// }

	// contents, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed read response: %s", err.Error())
	// }

	// var vkUserInfo models.VkUserInfo
	// var resp interface{}
	// err = json.Unmarshal([]byte(contents), &resp)
	// fmt.Println(resp, err)

	// if err != nil {
	// 	return nil, err
	// }

	// return &vkUserInfo, nil
	return nil, nil
}

func (s *Server) TelegramLogin(w http.ResponseWriter, p runtime.Producer) string {
	// redirect to
	// https://oauth.telegram.org/auth?bot_id=YOUR_BOT_ID&scope=YOUR_SCOPE&public_key=YOUR_PUBLIC_KEY&nonce=YOUR_NONCE&origin=http%3A%2F%2F<mydevorigin>

	nonce := randStringBytes(8)
	state := generateStateOauthCookie(w)
	//origin := settings.Settings.Origin

	authURL := settings.Settings.Telegram.AuthURL + "?bot_id=" + settings.Settings.Telegram.Token + "&public_key=" +
		settings.Settings.Telegram.PK + "&nonce=" + nonce + "&scope=all" + "&state=" + state + "&origin=https://dev.sheira.ru/oauth2/telegram"
	s.logger.Println(authURL)
	return authURL
}

func (s *Server) TelegramCallback(code string) (*models.VkUserInfo, error) {
	return getUserDataFromTelegram(code)
}

func randStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
