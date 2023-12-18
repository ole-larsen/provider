package settings

import (
	"os"

	log "github.com/olelarssen/provider/service/logger"
	"github.com/spf13/viper"
)

var Settings = initSettings()

type Google struct {
	ClientID     string
	ClientSecret string
	Callback     string
}

type Yandex struct {
	ClientID     string
	ClientSecret string
	Callback     string
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
}

type settings struct {
	Logger log.Logger
	Domain string
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
	// ###################################################################

	logger.Println("load settings done √")
	return ss
}
