package config_test

import (
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/Varshi292/RoastWear/internal/utils"
	"testing"
	"time"
)

func TestSessionConfigValidation(t *testing.T) {
	if err := utils.InitializeValidator(); err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name       string
		cfg        config.SessionConfig
		shouldFail bool
	}

	tests := []testCase{
		{
			name: "Valid Configuration",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Invalid: Missing Key",
			cfg: config.SessionConfig{
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: true,
		},
		{
			name: "Invalid: Missing Path",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: true,
		},
		{
			name: "Invalid: Missing MaxAge",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: true,
		},
		{
			name: "Valid: Missing Domain",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: Missing Secure",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: Missing HttpOnly",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: Missing SameSite",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: Missing SessionOnly",
			cfg: config.SessionConfig{
				Key:      "secret-key",
				Path:     "/",
				MaxAge:   720 * time.Hour,
				Domain:   "",
				Secure:   false,
				HttpOnly: true,
				SameSite: "Lax",
			},
			shouldFail: false,
		},
		{
			name: "Invalid: Empty Key",
			cfg: config.SessionConfig{
				Key:         "",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: true,
		},
		{
			name: "Invalid: Empty Path",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: true,
		},
		{
			name: "Invalid: Empty SameSite",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: Secure True",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      true,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: Secure False",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: HttpOnly True",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: HttpOnly False",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    false,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Valid: SessionOnly True",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: true,
			},
			shouldFail: false,
		},
		{
			name: "Valid: SessionOnly False",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: false,
		},
		{
			name: "Invalid: Invalid Path",
			cfg: config.SessionConfig{
				Key:         "secret-key",
				Path:        "/invalid",
				MaxAge:      720 * time.Hour,
				Domain:      "",
				Secure:      false,
				HttpOnly:    true,
				SameSite:    "Lax",
				SessionOnly: false,
			},
			shouldFail: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := utils.Validate.Struct(tc.cfg)
			if tc.shouldFail && err == nil {
				t.Errorf("Expected validation to fail but it passed.")
			}
			if !tc.shouldFail && err != nil {
				t.Errorf("Expected validation to pass but it failed: %v", err)
			}
		})
	}
}
