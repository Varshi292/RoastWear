package config

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT" validate:"required"`
	StaticFilesPath string `mapstructure:"STATIC_FILES_PATH" validate:"required"`
	DBPath          string `mapstructure:"DB_PATH" validate:"required"`
}

var validate = validator.New()

func Load() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config reading failed: %w", err)
	}
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("config unmarshalling failed: %w", err)
	}
	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return config, nil
}
