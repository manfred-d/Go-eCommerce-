package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	ServerAddr string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return config, errors.New("config file not found")
		}
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Error unmarshaling config: %v\n", err)
		return
	}

	log.Print("Loaded config")
	return
}
