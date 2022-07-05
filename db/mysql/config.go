package mysql

import (
	"fmt"
	"log"
	"os"
)

type Config interface {
	Dsn() string
	DBName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
	dsn    string
}

func NewConfig() (Config, error) {

	var cfg config
	var err error

	cfg.dbHost = os.Getenv("MYSQL_HOST")
	cfg.dbUser = os.Getenv("MYSQL_USER")
	cfg.dbPass = os.Getenv("MYSQL_PASS")
	cfg.dbName = os.Getenv("MYSQL_NAME")
	cfg.dbPort = os.Getenv("MYSQL_PORT")
	if err != nil {
		log.Print(fmt.Sprintf("error: cannot convert mysql port from string to int: %s", err.Error()))
		return nil, err
	}

	cfg.dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=1", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)

	return &cfg, nil
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DBName() string {
	return c.dbName
}
