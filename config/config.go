package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Conf Config

type Config struct {
	DB map[string]DBConfig `yaml:"DB"`
}

type DBConfig struct {
	Host string `yaml:"DB_HOST"`
	Port string `yaml:"DB_PORT"`
	User string `yaml:"DB_USER"`
	Pass string `yaml:"DB_PASS"`
	Name string `yaml:"DB_NAME"`
}

func loadConfig() Config {

	var conf Config

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("err while paring Yaml config file : ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal("err while parsing Yaml config file : ", err)
		os.Exit(1)
	}

	for env, dbConfig := range conf.DB {
		fmt.Printf("Environment: %s, DB Config: %+v\n", env, dbConfig)
	}

	return conf
}

func GetConfig() Config {
	if len(Conf.DB) == 0 {
		Conf = loadConfig()
	}
	return Conf
}
