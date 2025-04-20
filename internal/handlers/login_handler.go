package handlers

import (
	"errors"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

// LoginHandler ...
type LoginHandler struct {
	authService *services.AuthService
	sessionRepo *repositories.SessionRepository
}

// NewLoginHandler ...
func NewLoginHandler(auth *services.AuthService, sessionRepo *repositories.SessionRepository) *LoginHandler {
	return &LoginHandler{
		authService: auth,
		sessionRepo: sessionRepo,
	}
}

// UserLogin handles the login process for a user.
func (handler *LoginHandler) UserLogin(c *fiber.Ctx) error {
	var request models.UserLoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request format!",
			"details": err.Error(),
		})
	}

	if request.Username == "" || request.Password == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "All fields are required!",
			"details": "username and password are required",
		})
	}

	if err := handler.authService.AuthenticateUser(&request); err != nil {
		if errors.Is(utils.ErrInvalidCredentials, err) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials. Please check that you have provided the correct username and password.",
				"details": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error has occurred. Please contact support.",
			"details": err.Error(),
		})
	}

	// Create a new session
	sess, err := sessions.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create session.",
			"details": err.Error(),
		})
	}
	sess.Set("username", request.Username)
	sess.Set("loginTime", time.Now().Unix())

	// Wraps session in session model for GORM management and persistent storage
	modelSession := &models.Session{Session: sess, SessionKey: sess.ID()}

	// Stores session in database
	if err := handler.sessionRepo.CreateSession(modelSession); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to store session",
			"details": err.Error(),
		})
	}

	// Saves session as client-side cookie
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save session",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Login successful!",
		"success":    true,
		"session_id": modelSession.SessionKey,
	})
}
