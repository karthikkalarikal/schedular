package config

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"GRPC_PORT"`
}

func LoadConfig() (Config, error) {
	var config Config
	var envs = []string{"GRPC_PORT"}

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			log.Println("error in config", err)
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Println("error in config 2: ", err)
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		log.Println("error in config 3: ", err)
		return config, err
	}

	return config, nil
}
