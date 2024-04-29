package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Dialect(dsn string) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: dsn,
	})
}
