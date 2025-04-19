package bootstrap

import (
	"fmt"
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/spf13/viper"
)

func LoadSessionConfig(v *viper.Viper) (*config.SessionConfig, error) {
	// Default values
	v.SetDefault("SESSION_KEY", "secret-key")
	v.SetDefault("SESSION_PATH", "/")
	v.SetDefault("SESSION_MAX_AGE", "720h")

	var cfg config.SessionConfig

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
