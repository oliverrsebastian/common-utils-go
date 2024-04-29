package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dialect(dsn string) gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
}
