package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type YmlConfig struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	DBName string `yaml:"name"`
}

type Config struct {
	Host   string
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
		host := checkEnv("HOST")
		port := checkEnv("PORT")
		dbName := checkEnv("DB_NAME")
		return &Config{
			Host:   host,
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
