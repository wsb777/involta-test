package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type YmlConfig struct {
	DBHost string `yaml:"db_host"`
	DBPort string `yaml:"db_port"`
	DBName string `yaml:"db_name"`
}

type Config struct {
	DBHost string
	DBPort string
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
		host := checkEnv("DB_HOST")
		port := checkEnv("DB_PORT")
		dbName := checkEnv("DB_NAME")
		return &Config{
			DBHost: host,
			DBPort: port,
			DBName: dbName,
		}, nil
	}

	var ymlConfig YmlConfig
	err = yaml.Unmarshal(data, &ymlConfig)

	if err != nil {
		panic(err)
	}

	return &Config{
		DBHost: ymlConfig.DBHost,
		DBPort: ymlConfig.DBPort,
		DBName: ymlConfig.DBName,
	}, nil
}
