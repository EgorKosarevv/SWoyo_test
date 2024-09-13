package store

import (
	"sync"
)

// MemoryStore реализует интерфейс Store с использованием памяти для хранения данных
type MemoryStore struct {
	mu   sync.RWMutex
	urls map[string]string
}

// NewMemoryStore создает новый экземпляр MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urls: make(map[string]string),
	}
}

// Save сохраняет короткий URL и оригинальный URL
func (s *MemoryStore) Save(shortURL, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortURL] = originalURL
	return nil
}

// GetOriginalURL возвращает оригинальный URL по короткому URL
func (s *MemoryStore) GetOriginalURL(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	originalURL, exists := s.urls[shortURL]
	if !exists {
		return "", ErrNotFound
	}
	return originalURL, nil
}

// Exists проверяет, существует ли короткий URL
func (s *MemoryStore) Exists(shortURL string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.urls[shortURL]
	return exists, nil
}
