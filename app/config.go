package app

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Test string `yaml:"test" env-default:"test"`
}

var AppConfig Config

func ConfigInit() {
	if err := cleanenv.ReadConfig("config/config.yml", &AppConfig); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
}
