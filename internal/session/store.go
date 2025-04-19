package session

import (
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"time"
)

var Store *session.Store

func InitializeSessionStore(config *config.SessionConfig) {
	Store = session.New(session.Config{
		KeyLookup:         "cookie:" + config.Key,
		CookiePath:        config.Path,
		Expiration:        time.Duration(config.MaxAge),
		CookieDomain:      config.Domain,
		CookieSecure:      config.Secure,
		CookieHTTPOnly:    config.HttpOnly,
		CookieSameSite:    config.SameSite,
		CookieSessionOnly: config.SessionOnly,
	})
	log.Printf("✅ Session store initialized successfully.")
}
