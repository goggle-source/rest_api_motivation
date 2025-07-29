package main

import (
	"io"
	"os"

	"github.com/rest_api_motivation/internal/config"
	"github.com/rest_api_motivation/internal/database"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()
	log := NewLogger(cfg.Env)
	_ = log
	db := database.Init(*cfg)
	_ = db
	log.Info("good")
	//TODO: router(gin)
}

func NewLogger(env string) *logrus.Logger {
	var log logrus.Logger
	log.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error file")
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	switch env {
	case "Local":
		log.SetLevel(logrus.InfoLevel)

	case "Debug":
		log.SetLevel(logrus.DebugLevel)

	case "Prod":
		log.SetLevel(logrus.InfoLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	return &log
}
