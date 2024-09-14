package config

import (
	"github.com/spf13/viper"
)

const (
	defaultMongoDBURI = "mongodb://localhost:27017"
	defaultPort       = 8080
)

type AppConfig struct {
	App struct {
		Env  string
		Port uint16
	}

	DB struct {
		URI string
	}
}

var cfg *AppConfig

func Config() *AppConfig {
	if cfg == nil {
		loadConfig()
	}

	return cfg
}

func loadConfig() {

	// set defaults and override from env
	setDefaultValues()

	cfg = &AppConfig{}

	// App.
	cfg.App.Port = viper.GetUint16("APP_PORT")
	//cfg.App.Env = viper.GetString("APP_ENV")

	// Gin.
	//cfg.Gin.Mode = viper.GetString("GIN_MODE")

	//db
	cfg.DB.URI = viper.GetString("DB_URI")
}

func setDefaultValues() {
	viper.SetDefault("APP_PORT", defaultPort)
	viper.SetDefault("DB_URI", defaultMongoDBURI)
}
