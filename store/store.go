package store

import "errors"

// Error messages
var (
	ErrNotFound = errors.New("not found")
)

// Store интерфейс для работы с хранилищем URL
type Store interface {
	Save(shortURL, originalURL string) error
	GetOriginalURL(shortURL string) (string, error)
	Exists(shortURL string) (bool, error)
}
