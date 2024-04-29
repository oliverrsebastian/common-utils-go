package gcs

var config *Config

type Config struct {
	ProjectID         string `envconfig:"PROJECT_ID" required:"false"`
	ProjectKey        string `envconfig:"PROJECT_KEY" required:"false"`
	ProjectBucketName string `envconfig:"BUCKET_NAME" required:"false" DEFAULT:"bucket-name"`
	DefaultUrl        string `envconfig:"DEFAULT_URL" required:"false" default:"https://www.google.com/favicon.ico"`
	IsEnabled         bool   `envconfig:"ENABLED" required:"false" default:"false"`
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
