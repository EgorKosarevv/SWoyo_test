package store

import (
	"sync"
)

type MemoryStore struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urls: make(map[string]string),
	}
}

func (s *MemoryStore) Save(shortURL, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortURL] = originalURL
	return nil
}

func (s *MemoryStore) GetOriginalURL(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	originalURL, exists := s.urls[shortURL]
	if !exists {
		return "", ErrNotFound
	}
	return originalURL, nil
}

func (s *MemoryStore) Exists(shortURL string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.urls[shortURL]
	return exists, nil
}
