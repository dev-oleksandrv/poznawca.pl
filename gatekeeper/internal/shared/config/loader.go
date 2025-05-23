package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log/slog"
)

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		slog.Warn("cannot read config", "error", err)
	}

	viper.AutomaticEnv()

	config := new(Config)
	config.AppProxy.Port = viper.GetInt("APP_PROXY_PORT")
	config.AppProxy.Env = viper.GetString("APP_PROXY_ENV")
	config.BackofficeProxy.Port = viper.GetInt("BACKOFFICE_PROXY_PORT")
	config.BackofficeProxy.Env = viper.GetString("BACKOFFICE_PROXY_ENV")
	config.Postgres.Url = viper.GetString("POSTGRES_URL")
	config.Postgres.AutoMigrateEnabled = viper.GetBool("POSTGRES_AUTO_MIGRATE_ENABLED")
	config.WebClient.Url = viper.GetString("WEB_CLIENT_URL")

	return config, nil
}
