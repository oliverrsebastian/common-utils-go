package database

type Config struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     string `envconfig:"PORT" required:"true"`
	Username string `envconfig:"USERNAME" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Database string `envconfig:"DATABASE_NAME" required:"true"`

	Dialect string `envconfig:"DIALECT" required:"false" default:"postgresql"`
}
