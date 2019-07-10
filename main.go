package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/koind/shortener-sample/httpserver"
	"github.com/koind/shortener-sample/myshortener"
	"github.com/koind/shortener-sample/service"
)

func main() {

	httpPort, err := strconv.Atoi(os.Getenv("SHORTENER_PORT"))
	if err != nil {
		panic(fmt.Sprint("SHORTENER_PORT not defined"))
	}

	shortenerHost := os.Getenv("SHORTENER_HOST")
	if shortenerHost == "" {
		panic(fmt.Sprint("SHORTENER_HOST not defined"))
	}

	shortenerAddress := fmt.Sprintf("%s:%d", shortenerHost, httpPort)

	shortener := myshortener.NewMyShortener()
	shortenerService := service.NewShortenerService(shortener, shortenerAddress)
	hs := httpserver.NewHTTPServer(shortenerService, httpPort)

	log.Fatal(hs.Start())
}
