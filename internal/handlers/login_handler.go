package handlers

import (
	"errors"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// LoginHandler ...
type LoginHandler struct {
	authService    *services.AuthService
	sessionService *repositories.SessionRepository
}

// NewLoginHandler ...
func NewLoginHandler(auth *services.AuthService, session *repositories.SessionRepository) *LoginHandler {
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
	sess, err := handler.authService.LoginUser(&request, c)
	if err != nil {
		if errors.Is(utils.ErrInvalidCredentials, err) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials. Please ensure you have provided the correct username and password.",
				"details": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error has occurred. Please contact support.",
			"details": err.Error(),
		})
	}

	id := sess.Session.ID()
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save session",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Login successful!",
		"success":    true,
		"session_id": id,
	})
}
