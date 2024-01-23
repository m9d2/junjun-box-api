package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Conf = Config{}
)

type Config struct {
	Sqlite struct {
		DataPath string `yaml:"data-path"`
	}
	Wx struct {
		AppID     string `yaml:"appid"`
		AppSecret string `yaml:"app-secret"`
	}
	Cos struct {
		RawUrl    string `yaml:"raw-url"`
		SecretID  string `yaml:"secret-id"`
		SecretKey string `yaml:"secret-key"`
	}
	Mysql struct {
		Dsn string `yaml:"dsn"`
	}
}

func InitConfig() {
	file, _ := os.ReadFile("./config.yaml")
	err := yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
