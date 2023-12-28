package settings

import (
	"os"

	log "github.com/olelarssen/provider/service/logger"
	"github.com/spf13/viper"
)

var Settings = initSettings()

type Telegram struct {
	Token   string
	PK      string
	AuthURL string
}
type Google struct {
	ClientID     string
	ClientSecret string
	Callback     string
	Redirect     string
}

type Yandex struct {
	ClientID     string
	ClientSecret string
	Callback     string
	Redirect     string
}

type Vk struct {
	ClientID     string
	ClientSecret string
	Callback     string
}

type Auth struct {
	Google
	Yandex
	Vk
	Telegram
}

type settings struct {
	Logger log.Logger
	Domain string
	Origin string
	Auth
}

func initSettings() settings {
	var ss settings
	logger := log.NewLogger()

	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		logger.Errorln("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string

	ss.Logger = logger
	domain, ok := viper.Get("DOMAIN").(string)
	if !ok {
		domain = os.Getenv("DOMAIN")
	}
	ss.Domain = domain

	// ###################################################################

	ss.Auth.Google.ClientID, ok = viper.Get("NEXT_PUBLIC_GOOGLE_CLIENT_ID").(string)

	if !ok {
		ss.Auth.Google.ClientID = os.Getenv("NEXT_PUBLIC_GOOGLE_CLIENT_ID")
	}

	ss.Auth.Google.ClientSecret, ok = viper.Get("NEXT_PUBLIC_GOOGLE_CLIENT_SECRET").(string)

	if !ok {
		ss.Auth.Google.ClientSecret = os.Getenv("NEXT_PUBLIC_GOOGLE_CLIENT_SECRET")
	}

	ss.Auth.Google.Callback, ok = viper.Get("NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK").(string)

	if !ok {
		ss.Auth.Google.Callback = os.Getenv("NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK")
	}

	ss.Auth.Google.Redirect, ok = viper.Get("NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT").(string)

	if !ok {
		ss.Auth.Google.Redirect = os.Getenv("NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT")
	}

	// ###################################################################

	ss.Auth.Yandex.ClientID, ok = viper.Get("NEXT_PUBLIC_YANDEX_CLIENT_ID").(string)

	if !ok {
		ss.Auth.Yandex.ClientID = os.Getenv("NEXT_PUBLIC_YANDEX_CLIENT_ID")
	}

	ss.Auth.Yandex.ClientSecret, ok = viper.Get("NEXT_PUBLIC_YANDEX_CLIENT_SECRET").(string)

	if !ok {
		ss.Auth.Yandex.ClientSecret = os.Getenv("NEXT_PUBLIC_YANDEX_CLIENT_SECRET")
	}

	ss.Auth.Yandex.Callback, ok = viper.Get("NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK").(string)

	if !ok {
		ss.Auth.Yandex.Callback = os.Getenv("NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK")
	}

	ss.Auth.Yandex.Redirect, ok = viper.Get("NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT").(string)

	if !ok {
		ss.Auth.Yandex.Redirect = os.Getenv("NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT")
	}
	// ###################################################################

	ss.Auth.Vk.ClientID, ok = viper.Get("NEXT_PUBLIC_VK_CLIENT_ID").(string)

	if !ok {
		ss.Auth.Vk.ClientID = os.Getenv("NEXT_PUBLIC_VK_CLIENT_ID")
	}

	ss.Auth.Vk.ClientSecret, ok = viper.Get("NEXT_PUBLIC_VK_CLIENT_SECRET").(string)

	if !ok {
		ss.Auth.Vk.ClientSecret = os.Getenv("NEXT_PUBLIC_VK_CLIENT_SECRET")
	}

	ss.Auth.Vk.Callback, ok = viper.Get("NEXT_PUBLIC_VK_CLIENT_CALLBACK").(string)

	if !ok {
		ss.Auth.Vk.Callback = os.Getenv("NEXT_PUBLIC_VK_CLIENT_CALLBACK")
	}

	ss.Auth.Telegram.Token, ok = viper.Get("NEXT_PUBLIC_TELEGRAM_BOT_KEY").(string)
	if !ok {
		ss.Auth.Telegram.Token = os.Getenv("NEXT_PUBLIC_TELEGRAM_BOT_KEY")
	}

	ss.Auth.Telegram.PK, ok = viper.Get("NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY").(string)
	if !ok {
		ss.Auth.Telegram.PK = os.Getenv("NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY")
	}

	ss.Auth.Telegram.AuthURL, ok = viper.Get("NEXT_PUBLIC_TELEGRAM_AUTH_URL").(string)
	if !ok {
		ss.Auth.Telegram.AuthURL = os.Getenv("NEXT_PUBLIC_TELEGRAM_AUTH_URL")
	}

	ss.Origin, ok = viper.Get("NEXT_PUBLIC_URL").(string)
	if !ok {
		ss.Origin = os.Getenv("NEXT_PUBLIC_URL")
	}
	// ###################################################################

	logger.Println(ss.Origin)
	logger.Println("load settings done √")
	return ss
}
