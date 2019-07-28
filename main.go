package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/koind/shortener-servis/config"
	"github.com/koind/shortener-servis/httpserver"
	"github.com/koind/shortener-servis/myshortener"
	"github.com/koind/shortener-servis/mystats"
	"github.com/koind/shortener-servis/service"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logger.Sync()

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	shortenerAddress := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	shortener := myshortener.NewMyShortener()
	stats := mystats.NewStats(logger)
	shortenerService := service.NewShortenerService(shortener, stats, logger, shortenerAddress)
	hs := httpserver.NewHTTPServer(shortenerService, cfg.Port)

	logger.Error("Error starting app", zap.Error(hs.Start()))
}
