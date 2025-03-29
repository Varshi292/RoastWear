package config

import (
	"net/http"
)

type SessionConfig struct {
	Key         string        `mapstructure:"SESSION_KEY" validate:"required"`
	Path        string        `mapstructure:"SESSION_PATH" validate:"required"`
	MaxAge      int           `mapstructure:"SESSION_MAX_AGE" validate:"required"`
	Domain      string        `mapstructure:"SESSION_DOMAIN"`
	Secure      bool          `mapstructure:"SESSION_SECURE"`
	HttpOnly    bool          `mapstructure:"SESSION_HTTP_ONLY"`
	Partitioned bool          `mapstructure:"SESSION_PARTITIONED"`
	SameSite    http.SameSite `mapstructure:"SESSION_SAME_SITE"`
}
