package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect создает соединение с базой данных и возвращает указатель на sql.DB
func Connect() (*sql.DB, error) {
	// Строка подключения
	connStr := "user=postgres dbname=myurlshort sslmode=disable password=1234512345"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Проверка соединения
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	return db, nil
}
