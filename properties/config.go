package properties

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Properties *Config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Jwt      JwtConfig      `yaml:"jwt"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
}

func InitProperties() {
	data, err := os.ReadFile("application_properties.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &Properties)
	if err != nil {
		log.Fatal(err)
	}
}
