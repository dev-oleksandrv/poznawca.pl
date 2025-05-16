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

	config.AdminProxy.Port = viper.GetInt("ADMIN_PROXY_PORT")
	config.AdminProxy.Env = viper.GetString("ADMIN_PROXY_ENV")

	config.OpenAI.APIKey = viper.GetString("OPENAI_API_KEY")
	config.OpenAI.InterviewAssistantID = viper.GetString("OPENAI_INTERVIEW_ASSISTANT_ID")

	config.Postgres.Url = viper.GetString("POSTGRES_URL")
	config.Postgres.AutoMigrateEnabled = viper.GetBool("POSTGRES_AUTO_MIGRATE_ENABLED")

	config.Redis.Addr = viper.GetString("REDIS_ADDR")
	config.Redis.Password = viper.GetString("REDIS_PASSWORD")
	config.Redis.DB = viper.GetInt("REDIS_DB")

	return config, nil
}
