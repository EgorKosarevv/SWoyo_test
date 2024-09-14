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
		Length      int `yaml:"length"`
		MaxLength   int `yaml:"maxLength"`
		MaxAttempts int `yaml:"maxAttempts"`
	} `yaml:"url"`
}

var Cfg Config

func LoadConfig() {
	file, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("Error reading the configuration file: %v", err)
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalf("Error parsing the configuration file: %v", err)
	}
}
