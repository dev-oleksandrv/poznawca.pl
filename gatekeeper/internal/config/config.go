package config

type AppProxyConfig struct {
	Port int
	Env  string
}

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
	AppProxy   AppProxyConfig
	AdminProxy AdminProxyConfig
	OpenAI     OpenAIConfig
	Postgres   PostgresConfig
}
