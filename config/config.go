package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"db"`
	URL struct {
		Length int `yaml:"length"`
	} `yaml:"url"`
}

var Cfg Config

func LoadConfig() {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Ошибка чтения файла конфигурации: %v", err)
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalf("Ошибка разбора файла конфигурации: %v", err)
	}

	// Для отладки: выведите загруженную конфигурацию
	log.Printf("Загружена конфигурация: %+v", Cfg)
}
