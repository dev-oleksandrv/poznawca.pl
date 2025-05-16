package config

type AdminProxyConfig struct {
	Port int
	Env  string
}

type OpenAIConfig struct {
	APIKey               string
	InterviewAssistantID string
}

type PostgresConfig struct {
	Url                string
	AutoMigrateEnabled bool
}

type Config struct {
	AdminProxy AdminProxyConfig
	OpenAI     OpenAIConfig
	Postgres   PostgresConfig
}
