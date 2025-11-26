package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-vayu/vayu/internal/config"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func InitDB() (*xorm.Engine, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DatabaseHost.GetString(),
		config.DatabasePort.GetInt(),
		config.DatabaseUser.GetString(),
		config.DatabasePassword.GetString(),
		config.DatabaseDatabase.GetString(),
		config.DatabaseSslMode.GetString(),
	)

	engine, err := xorm.NewEngine("postgres", connStr)
	if err != nil {
		return nil, err
	}

	engine.SetMaxIdleConns(config.DatabaseMaxIdleConnections.GetInt())
	engine.SetMaxOpenConns(config.DatabaseMaxOpenConnections.GetInt())
	maxLifetime, err := time.ParseDuration(strconv.Itoa(config.DatabaseMaxConnectionLifetime.GetInt()) + "ms")
	if err != nil {
		return nil, err
	}
	engine.SetConnMaxLifetime(maxLifetime)

	return engine, nil
}
