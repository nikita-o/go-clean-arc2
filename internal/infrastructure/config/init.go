package config

import (
	"github.com/spf13/viper"
	"log"
)

func NewConfig() *Config {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка при чтении конфигурации: %v", err)
	}

	// Привязываем значения к структуре
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Не удалось привязать конфиг к структуре: %v", err)
	}

	return &config
}
