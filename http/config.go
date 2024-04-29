package http

var config *Config

type Config struct {
	Port   int  `envconfig:"PORT"`
	UseTLS bool `envconfig:"USE_TLS" default:"false"`
	Secure bool `envconfig:"SECURED" default:"false"`
	CORS   bool `envconfig:"CORS_ENABLED" default:"false"`
	CSRF   bool `envconfig:"CSRF_ENABLED" default:"false"`
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
