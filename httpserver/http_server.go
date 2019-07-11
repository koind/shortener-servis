package httpserver

import (
	"fmt"
	"github.com/koind/shortener-servis/service"
	"net/http"

	"github.com/gorilla/mux"
)

// HttpServer represents transport layer of out app
type HttpServer struct {
	httpPort int
	router   http.Handler
	s        *service.ShortenerService
	stats    *service.StatsService
}

// Start fires up the http server
func (s *HttpServer) Start() error {
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.httpPort), s.router)
}

// NewHTTPServer returns http server that wraps shortener business logic
func NewHTTPServer(ss *service.ShortenerService, stats *service.StatsService, port int) *HttpServer {

	r := mux.NewRouter()
	hs := HttpServer{router: r, httpPort: port, s: ss, stats: stats}

	r.HandleFunc("/{shortened}", ss.ResolverHandle)
	r.HandleFunc("/", ss.ResolverHandle)

	http.Handle("/", r)
	http.HandleFunc("/stats", ss.StatsHandle)

	return &hs
}
