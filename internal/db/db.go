package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-vayu/vayu/internal/config"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func InitDB() (*xorm.Engine, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.DatabaseUser.GetString(),
		config.DatabasePassword.GetString(),
		config.DatabaseHost.GetString(),
		config.DatabasePort.GetInt(),
		config.DatabaseDatabase.GetString(),
		config.DatabaseSslMode.GetString(),
	)

	engine, err := xorm.NewEngine("postgres", connStr)
	if err != nil {
		return nil, err
	}

	engine.SetMapper(names.LintGonicMapper)

	engine.SetMaxIdleConns(config.DatabaseMaxIdleConnections.GetInt())
	engine.SetMaxOpenConns(config.DatabaseMaxOpenConnections.GetInt())
	maxLifetime, err := time.ParseDuration(strconv.Itoa(config.DatabaseMaxConnectionLifetime.GetInt()) + "ms")
	if err != nil {
		return nil, err
	}
	engine.SetConnMaxLifetime(maxLifetime)

	return engine, nil
}
