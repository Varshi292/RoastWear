package handlers

import (
	"errors"

	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// LoginHandler ...
type LoginHandler struct {
	authService    *services.AuthService
	sessionService *services.SessionService
}

// NewLoginHandler ...
func NewLoginHandler(auth *services.AuthService, session *services.SessionService) *LoginHandler {
	return &LoginHandler{
		authService:    auth,
		sessionService: session,
	}
}

// UserLogin handles the login process for a user.
func (handler *LoginHandler) UserLogin(c *fiber.Ctx) error {
	var request models.UserLoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid request format",
			"details": err.Error(),
		})
	}

	if request.Username == "" || request.Password == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "validation error",
			"details": "username and password are required",
		})
	}

	if err := handler.authService.LoginUser(&request, c); err != nil {
		if errors.Is(utils.ErrInvalidCredentials, err) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "invalid credentials",
				"details": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "internal server error",
			"details": err.Error(),
		})
	}

	// ✅ Generate session ID and store it
	sessionID := uuid.New().String()
	session := &models.Session{
		Username:  request.Username,
		SessionID: sessionID,
	}

	if err := handler.sessionService.CreateSession(session); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to create session",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "login successful",
		"session_id": sessionID,
	})
}
