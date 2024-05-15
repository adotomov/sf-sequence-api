package config

import (
	"database/sql"
	"time"
)

type httpConfig struct {
	Addr              string        `env:"HTTP_ADDR"`
	ReadTimeout       time.Duration `env:"HTTP_READ_TIMEOUT"`
	ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT"`
	WriteTimeout      time.Duration `env:"HTTP_WRITE_TIMEOUT"`
	IdleTimeout       time.Duration `env:"HTTP_IDLE_TIMEOUT"`
	MaxHeaderBytes    int           `env:"HTTP_MAX_HEADER_BYTES"`
	GracefulShutdown  time.Duration `env:"HTTP_GRACEFUL_SHUTDOWN,default=10s"`
}

type DBConfig struct {
	DSN             string        `env:"SQL_DSN"`
	Driver          string        `env:"SQL_DRIVER"`
	MaxIdleConns    int           `env:"SQL_MAX_IDLE_CONNS,default=2"`
	MaxOpenConns    int           `env:"SQL_MAX_OPEN_CONNS"`
	ConnMaxIdleTime time.Duration `env:"SQL_CONN_MAX_IDLE_TIME"`
	ConnMaxLifetime time.Duration `env:"SQL_CONN_MAX_LIFETIME"`
}

type AppConfig struct {
	Env        string `env:"DEPLOYMENT_ENVIRONMENT"`
	HttpServer *httpConfig
	DBConfig   DBConfig
}

type AppDependencies struct {
	DB *sql.DB
}
