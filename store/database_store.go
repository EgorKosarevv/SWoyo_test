package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBStore struct {
	db *sql.DB
}

func NewDBStore(db *sql.DB) *DBStore {
	return &DBStore{db: db}
}

func (s *DBStore) Save(shortURL, originalURL string) error {
	_, err := s.db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortURL, originalURL)
	return err
}

func (s *DBStore) GetOriginalURL(shortURL string) (string, error) {
	var originalURL string
	err := s.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("query error: %w", err)
	}
	return originalURL, nil
}

func (s *DBStore) Exists(shortURL string) (bool, error) {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM urls WHERE short_url = $1)", shortURL).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("query error: %w", err)
	}
	return exists, nil
}
