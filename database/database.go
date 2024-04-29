package database

import (
	"common-utils-go/database/mysql"
	"common-utils-go/database/postgresql"
	"context"
	"fmt"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type Database interface {
	GetDB(ctx context.Context) Database
	startTx(ctx context.Context) *gorm.DB
	commit(tx *gorm.DB) *gorm.DB
	rollback(tx *gorm.DB) *gorm.DB
}

type queries struct {
	DB *gorm.DB
}

func New(config *Config) *queries {
	dialect := getDialect(config)
	gormDB, err := gorm.Open(dialect, &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}

	// recommended starting point for optimized db performance.
	// refer to here: https://www.alexedwards.net/blog/configuring-sqldb
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return &queries{
		DB: gormDB,
	}
}

func WithTransaction(ctx context.Context, db Database, fn func(ctx context.Context) error) error {
	tx := db.startTx(ctx)
	ctx = context.WithValue(ctx, TransactionCtxKey, tx)

	if err := fn(ctx); err != nil {
		db.rollback(tx)
		return err
	}
	db.commit(tx)
	return nil
}

func (db *queries) GetDB(ctx context.Context) Database {
	return &queries{
		DB: db.DB.WithContext(ctx),
	}
}

func (db *queries) startTx(ctx context.Context) *gorm.DB {
	return db.DB.WithContext(ctx).Begin()
}

func (db *queries) commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (db *queries) rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func getDialect(cfg *Config) gorm.Dialector {
	dsn := createDSN(cfg, false)
	switch cfg.Dialect {
	case "mysql":
		return mysql.Dialect(dsn)
	default:
		return postgresql.Dialect(dsn)
	}
}

func createDSN(cfg *Config, escaped bool) string {
	if escaped == true {
		authStr := fmt.Sprintf("%s:%s", url.QueryEscape(cfg.Username), url.QueryEscape(cfg.Password))
		dbConnStr := fmt.Sprintf("%s:%s/%s", cfg.Host, cfg.Port, cfg.Database)
		return fmt.Sprintf("postgres://%s@%s?sslmode=disable", authStr, dbConnStr)
	} else {
		authStr := fmt.Sprintf("%s:%s", cfg.Username, cfg.Password)
		dbConnStr := fmt.Sprintf("%s:%s/%s", cfg.Host, cfg.Port, cfg.Database)
		return fmt.Sprintf("postgresql://%s@%s?sslmode=disable", authStr, dbConnStr)
	}
}
