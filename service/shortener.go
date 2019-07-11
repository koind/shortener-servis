package service

import (
	"bytes"
	"fmt"
	"github.com/koind/shortener-servis/stats"
	"net/http"

	"github.com/gorilla/mux"
)

// Shortener implements business logic for the shortener service
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type ShortenerService struct {
	Shortener
	*StatsService
	address string
}

// NewShortenerService creates a new shortener service
func NewShortenerService(shortener Shortener, stats *StatsService, address string) *ShortenerService {
	return &ShortenerService{shortener, stats, address}
}

func (ss *ShortenerService) ResolverHandle(w http.ResponseWriter, r *http.Request) {
	ss.Stats.Add(r.URL.String())

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		if url, ok := vars["shortened"]; ok {
			shortUrl := ss.Shortener.Resolve(string(url))
			http.Redirect(w, r, shortUrl, http.StatusSeeOther)
		} else {
			w.WriteHeader(404)
		}
	case "POST":
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		shortened := ss.address + "/" + ss.Shortener.Shorten(buf.String())
		w.Write([]byte(shortened))
	}
}

type StatsService struct {
	Stats *stats.Stats
}

func (s *StatsService) StatsHandle(w http.ResponseWriter, r *http.Request) {
	s.Stats.Add(r.URL.String())

	str := ""
	allStats := s.Stats.GetAll()
	if len(allStats) < 0 {
		fmt.Fprint(w, str)
	}

	for url, count := range allStats {
		str += fmt.Sprintf("Url: %s, Count: %d\n", url, count)
	}
	fmt.Fprint(w, str)
}

func NewStatsService(stats *stats.Stats) *StatsService {
	return &StatsService{stats}
}
