package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	DbHost     string `yaml:"db-host"`
	DbUsername string `yaml:"db-username"`
	DbPassword string `yaml:"db-password"`
	DbName     string `yaml:"db-name"`
}

var AppConfig Config

func Load() {
	if err := cleanenv.ReadConfig("config/config.yml", &AppConfig); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
}
