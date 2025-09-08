package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type YmlConfig struct {
	Port   string `yaml:"port"`
	DBName string `yaml:"name"`
}

type EnvConfig struct {
	Port   string
	DBName string
}

type Config struct {
	Port   string
	DBName string
}

func checkEnv(value string) string {
	s := os.Getenv(value)
	if s == "" {
		panic("Нету переменной: " + value)
	}
	return s
}

func NewConfig() (*Config, error) {
	data, err := os.ReadFile("config.yml")

	if err != nil {
		log.Print("Файл config.yml не найден, проверка .env")
		port := checkEnv("PORT")
		dbName := checkEnv("DB_NAME")
		return &Config{
			Port:   port,
			DBName: dbName,
		}, nil
	}

	var ymlConfig YmlConfig
	err = yaml.Unmarshal(data, &ymlConfig)

	if err != nil {
		panic(err)
	}

	return &Config{
		Port:   ymlConfig.Port,
		DBName: ymlConfig.DBName,
	}, nil
}
