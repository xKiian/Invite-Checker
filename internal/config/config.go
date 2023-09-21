package config

import (
	"os"

	"gopkg.in/yaml.v2"
	"checker/internal/logger"
)


var (
	Config = LoadConfig()
	Logger = logger.Logger
)

type ConfigS struct {
	MinBoosts 	int `yaml:"min_boosts"`
	MinMembers 	int `yaml:"min_members"`
	MinOnline 	int `yaml:"min_online"`
}

func LoadConfig() ConfigS {
	var config ConfigS
	file, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}


	return config
}