package bootstrap

import (
	"fmt"
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

var validate = validator.New()

func LoadAppConfig(v *viper.Viper) (*config.AppConfig, error) {
	// Default values
	v.SetDefault("PORT", "7777")
	v.SetDefault("DB_PATH", "./db/users.db")
	v.SetDefault("STATIC_FILES_PATH", "./frontend/build")

	var cfg config.AppConfig

	// Unmarshal configuration
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate configuration
	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}
	return &cfg, nil
}
