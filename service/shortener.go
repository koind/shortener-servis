package service

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/koind/shortener-servis/mystats"
	"go.uber.org/zap"
	"net/http"
)

// Shortener implements business logic for the shortener service
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type ShortenerService struct {
	Shortener
	stats   mystats.StatsInterface
	logger  *zap.Logger
	address string
}

// NewShortenerService creates a new shortener service
func NewShortenerService(
	shortener Shortener,
	stats mystats.StatsInterface,
	logger *zap.Logger,
	address string,
) *ShortenerService {
	return &ShortenerService{shortener, stats, logger, address}
}

func (ss *ShortenerService) ResolverHandle(w http.ResponseWriter, r *http.Request) {
	ss.stats.Add(r.URL.String())

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		if url, ok := vars["shortened"]; ok {
			shortUrl := ss.Shortener.Resolve(string(url))

			ss.logger.Info(
				"Url was found",
				zap.String("url", url),
				zap.String("shortUrl", shortUrl),
			)

			http.Redirect(w, r, shortUrl, http.StatusSeeOther)
		} else {
			ss.logger.Warn("Url not found", zap.String("url", url))

			w.WriteHeader(404)
		}
	case "POST":
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		shortened := ss.address + "/" + ss.Shortener.Shorten(buf.String())

		ss.logger.Info("Created new short url", zap.String("shortUrl", shortened))

		w.Write([]byte(shortened))
	}
}

func (ss *ShortenerService) StatsHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ss.stats.GetAll()))
}
