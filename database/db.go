package database

import (
	"SWOYO/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Connect создает соединение с базой данных
func Connect() (*sql.DB, error) {
	// Жестко заданное значение sslmode
	sslMode := "disable"

	// Вывод значений конфигурации для отладки
	fmt.Printf("Config values: User=%s, DBName=%s, SSLMode=%s, Password=%s\n",
		config.Cfg.DB.User, config.Cfg.DB.DBName, sslMode, config.Cfg.DB.Password)

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s",
		config.Cfg.DB.User, config.Cfg.DB.DBName, sslMode, config.Cfg.DB.Password)
	fmt.Printf("Connecting to database with connection string: %s\n", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	return db, nil
}
