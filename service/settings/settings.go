package settings

import (
	log "github.com/olelarssen/provider/service/logger"
	"github.com/spf13/viper"
)

var Settings = initSettings()

type Google struct {
	ClientID     string
	ClientSecret string
	Callback     string
}

type Auth struct {
	Google
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
		logger.Printf(domain)
	}
	ss.Domain = domain

	ss.Auth.ClientID = viper.Get("NEXT_PUBLIC_GOOGLE_CLIENT_ID").(string)
	ss.Auth.ClientSecret = viper.Get("NEXT_PUBLIC_GOOGLE_CLIENT_SECRET").(string)
	ss.Auth.Callback = viper.Get("NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK").(string)
	logger.Println("load settings done √")
	return ss
}
