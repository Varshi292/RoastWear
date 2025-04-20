package bootstrap

import (
	"fmt"
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

var validate = validator.New()

func loadAppConfig(v *viper.Viper) (*config.AppConfig, error) {
	// Default values
	v.SetDefault("BACKEND_PORT", "7777")
	v.SetDefault("FRONTEND_PORT", "3000")
	v.SetDefault("MODE", "development")
	v.SetDefault("DEVELOPMENT_DOMAIN", "http://localhost")
	v.SetDefault("PRODUCTION_DOMAIN", "https://www.roastwear.com")
	v.SetDefault("USER_DB_PATH", "./db/users.db")
	v.SetDefault("SESSION_DB_PATH", "./db/sessions.db")
	v.SetDefault("UPLOAD_DB_PATH", "./db/uploads.db")
	v.SetDefault("CART_DB_PATH", "./db/carts.db")

	v.SetDefault("STATIC_FILES_PATH", "./frontend/build")

	var cfg config.AppConfig

	// Determine mode
	mode := v.GetString("MODE")
	switch mode {
	case "development":
		cfg.Domain = v.GetString("DEVELOPMENT_DOMAIN")
	case "production":
		cfg.Domain = v.GetString("PRODUCTION_DOMAIN")
	default:
		return nil, fmt.Errorf("unrecognized mode '%s'", mode)
	}

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
