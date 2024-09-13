package models

import (
	"SWOYO/store"
	"math/rand"
	"time"
)

const urlLength = 7
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Генерация сокращенного URL
func GenerateShortURL(originalURL string, storage store.Store) (string, error) {
	var shortURL string
	for {
		shortURL = randomString(urlLength)
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

// Генерация случайной строки
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
