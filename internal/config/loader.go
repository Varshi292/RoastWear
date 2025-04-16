// Package config ...
package config

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

// LoadConfig ...
//
// Parameters:
//   - config: ...
//
// Returns:
//   - error: ...
func LoadConfig(config interface{}) error {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("config reading failed: %w", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("config unmarshalling failed: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	return nil
}
