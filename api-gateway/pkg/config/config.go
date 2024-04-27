package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	EmployeeGRPC string `mapstructure:"EMPLOYEE_GRPC_CONNECTION"`
}

func LoadConfig() (Config, error) {
	var config Config
	var envs = []string{"PORT", "EMPLOYEE_GRPC_CONNECTION"}

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
