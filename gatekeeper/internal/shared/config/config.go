package config

type BaseServerConfig struct {
	Port int
	Env  string
}

type AppProxyConfig struct {
	BaseServerConfig
}

type BackofficeProxyConfig struct {
	BaseServerConfig
}

type PostgresConfig struct {
	Url                string
	AutoMigrateEnabled bool
}

type WebClientConfig struct {
	Url string
}

type Config struct {
	AppProxy        AppProxyConfig
	BackofficeProxy BackofficeProxyConfig
	Postgres        PostgresConfig
	WebClient       WebClientConfig
}
