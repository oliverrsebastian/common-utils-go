package redis

type Config struct {
	Host     string `envconfig:"HOST" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
}
