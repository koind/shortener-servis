package mystats

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

type StatsInterface interface {
	Add(url string)
	Get(url string) int
	GetAll() string
}

type Stats struct {
	sync.RWMutex
	data   map[string]int
	logger *zap.Logger
}

func NewStats(logger *zap.Logger) *Stats {
	return &Stats{data: make(map[string]int), logger: logger}
}

func (s *Stats) Add(url string) {
	s.Lock()
	defer s.Unlock()

	s.data[url]++
	s.logger.Info("The url has been incremented", zap.String("url", url))
}

func (s *Stats) Get(url string) int {
	s.RLock()
	count := s.data[url]
	defer s.RUnlock()

	return count
}

func (s *Stats) GetAll() string {
	s.RLock()
	resultData := s.data
	defer s.RUnlock()

	buf := new(bytes.Buffer)
	for key, value := range resultData {
		fmt.Fprintf(buf, "%s=\"%s\"\n", key, value)
	}

	return buf.String()
}
