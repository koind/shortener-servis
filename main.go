package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/koind/shortener-servis/httpserver"
	"github.com/koind/shortener-servis/myshortener"
	"github.com/koind/shortener-servis/service"
	"github.com/koind/shortener-servis/stats"
	"github.com/sirupsen/logrus"
)

var config struct {
	Host string
	Port int
}

func init() {
	if _, err := toml.DecodeFile("config/testing/config.toml", &config); err != nil {
		logrus.Fatalln("Failed to load config", err)
		return
	}
}

func main() {
	shortenerAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)

	shortener := myshortener.NewMyShortener()
	statsStruct := stats.NewStats()
	statsService := service.NewStatsService(statsStruct)
	shortenerService := service.NewShortenerService(shortener, statsService, shortenerAddress)
	hs := httpserver.NewHTTPServer(shortenerService, statsService, config.Port)

	logrus.Fatalln(hs.Start())
}
