package models

import (
	"SWOYO/config"
	"SWOYO/store"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateShortURL(originalURL string, storage store.Store) (string, error) {
	var shortURL string
	for {
		shortURL = randomString(config.Cfg.URL.Length)
		exists, _ := storage.Exists(shortURL)
		if !exists {
			break
		}
	}

	if err := storage.Save(shortURL, originalURL); err != nil {
		return "", err
	}
	return shortURL, nil
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
