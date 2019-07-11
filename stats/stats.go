package stats

import "sync"

type Stats struct {
	sync.RWMutex
	data map[string]int
}

func NewStats() *Stats {
	return &Stats{data: make(map[string]int)}
}

func (s *Stats) Add(url string) {
	s.Lock()
	defer s.Unlock()

	s.data[url]++
}

func (s *Stats) Get(url string) int {
	s.RLock()
	count := s.data[url]
	defer s.RUnlock()

	return count
}

func (s *Stats) GetAll() map[string]int {
	return s.data
}
