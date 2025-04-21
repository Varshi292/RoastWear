package utils

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/gofiber/fiber/v2"
	"time"
)

func StartSession(c *fiber.Ctx, r *repositories.SessionRepository, username string) (string, error) {
	sess, err := sessions.Store.Get(c)
	if err != nil {
		return "", err
	}
	sess.Set("username", username)
	sess.Set("loginTime", time.Now().Unix())

	// Wraps session in session model for GORM management and persistent storage
	modelSession := &models.Session{Session: sess, SessionKey: sess.ID()}

	// Stores session in database
	if err := r.CreateSession(modelSession); err != nil {
		return "", err
	}

	// Saves session as client-side cookie
	if err := sess.Save(); err != nil {
		return "", err
	}
	return modelSession.SessionKey, nil
}
