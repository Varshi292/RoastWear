package config

import (
	"time"
)

type SessionConfig struct {
	Key         string        `mapstructure:"SESSION_KEY" validate:"required"`
	Path        string        `mapstructure:"SESSION_PATH" validate:"required"`
	MaxAge      time.Duration `mapstructure:"SESSION_MAX_AGE" validate:"required"`
	Domain      string        `mapstructure:"SESSION_DOMAIN"`
	Secure      bool          `mapstructure:"SESSION_SECURE"`
	HttpOnly    bool          `mapstructure:"SESSION_HTTP_ONLY"`
	SameSite    string        `mapstructure:"SESSION_SAME_SITE"`
	SessionOnly bool          `mapstructure:"SESSION_SESSION_ONLY"`
}
