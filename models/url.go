package models

import (
	"SWOYO/config"
	"SWOYO/store"
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Short URL generation
func GenerateShortURL(originalURL string, storage store.Store) (string, error) {
	maxLength := config.Cfg.URL.Length
	for {
		shortURL, err := generateUniqueURL(originalURL, storage, maxLength)
		if err != nil {
			return "", err
		}
		if shortURL != "" {
			return shortURL, nil
		}
		if maxLength < config.Cfg.URL.MaxLength {
			maxLength++
			config.Cfg.URL.Length = maxLength
		} else {
			return "", fmt.Errorf("cannot generate unique URL; maximum length reached")
		}
	}
}

// Generating a unique URL
func generateUniqueURL(originalURL string, storage store.Store, length int) (string, error) {
	for i := 0; i < config.Cfg.URL.MaxAttempts; i++ {
		shortURL := randomString(length)
		exists, _ := storage.Exists(shortURL)
		if !exists {
			if err := storage.Save(shortURL, originalURL); err != nil {
				return "", err
			}
			return shortURL, nil
		}
	}
	return "", nil
}

// Random string generation
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
