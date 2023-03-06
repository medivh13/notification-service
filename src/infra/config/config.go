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

type SlackConf struct {
	WebHookURL string
	Icon       string
	Emoji      string
	Channel    string
}

type SMTPConf struct {
	UserName string
	Password string
	Host     string
	Port     int
}

// Config ...
type Config struct {
	App  AppConf
	Http HttpConf
	Log  LogConf
	SMTP SMTPConf
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

		SMTP: smtp,
	}

	return config
}
