package mysql

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewConnection(cfg Config) (*sqlx.DB, error) {
	sqlDB, err := sqlx.Connect("mysql", cfg.Dsn())
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return sqlDB, nil
}
