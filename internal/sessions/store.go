package sessions

import (
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func InitializeSessionStore(config *config.SessionConfig) {
	Store = session.New(session.Config{
		KeyLookup:         "cookie:" + config.Key,
		CookiePath:        config.Path,
		Expiration:        config.MaxAge,
		CookieDomain:      config.Domain,
		CookieSecure:      config.Secure,
		CookieHTTPOnly:    config.HttpOnly,
		CookieSameSite:    config.SameSite,
		CookieSessionOnly: config.SessionOnly,
	})
}
