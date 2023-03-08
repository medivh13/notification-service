package config

import (
	"os"
	"strconv"
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type LogConf struct {
	Name string
}

type NatsConf struct {
	NatsHost   string
	NatsStatus string
}

type SMTPConf struct {
	UserName string
	Password string
	Host     string
	Port     int
}

type SqlDbConf struct {
	Host                   string
	Username               string
	Password               string
	Name                   string
	Port                   string
	SSLMode                string
	Schema                 string
	MaxOpenConn            int
	MaxIdleConn            int
	MaxIdleTimeConnSeconds int64
	MaxLifeTimeConnSeconds int64
}

// Config ...
type Config struct {
	App   AppConf
	Http  HttpConf
	Log   LogConf
	SMTP  SMTPConf
	SqlDb SqlDbConf
	Nats  NatsConf
}

// NewConfig ...
func Make() Config {
	app := AppConf{
		Environment: os.Getenv("APP_ENV"),
		Name:        os.Getenv("APP_NAME"),
	}

	http := HttpConf{
		Port:       os.Getenv("HTTP_PORT"),
		XRequestID: os.Getenv("HTTP_REQUEST_ID"),
	}

	log := LogConf{
		Name: os.Getenv("LOG_NAME"),
	}

	sqldb := SqlDbConf{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}

	nats := NatsConf{
		NatsHost:   os.Getenv("NATS_HOST"),
		NatsStatus: os.Getenv("NATS_STATUS"),
	}

	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	smtp := SMTPConf{
		UserName: os.Getenv("MAIL_USER"),
		Password: os.Getenv("MAIL_PASSWORD"),
		Host:     os.Getenv("MAIL_HOST"),
		Port:     port,
	}

	config := Config{
		App:  app,
		Http: http,
		Log:  log,

		SMTP:  smtp,
		SqlDb: sqldb,
		Nats:  nats,
	}

	return config
}
