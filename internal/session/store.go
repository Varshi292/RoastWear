package session

import (
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/gorilla/sessions"
	"log"
)

var Store *sessions.CookieStore

func InitializeSessionStore(config config.SessionConfig) {
	Store = sessions.NewCookieStore([]byte(config.Key))
	Store.Options = &sessions.Options{
		Path:        config.Path,
		MaxAge:      config.MaxAge,
		Domain:      config.Domain,
		Secure:      config.Secure,
		HttpOnly:    config.HttpOnly,
		Partitioned: config.Partitioned,
		SameSite:    config.SameSite,
	}
	log.Printf("✅ Session store initialized successfully.")
}
