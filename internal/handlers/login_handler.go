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
// @Summary User login
// @Description This endpoint handles the login process for a user by verifying credentials and starting a session.
// @Tags Auth
// @Accept json
// @Produce json
// @Param userLoginRequest body models.UserLoginRequest true "User login details"
// @Success 200 {object} models.LoginSuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /login [post]

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
	sessionKey, err := utils.StartSession(c, handler.sessionRepo, request.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error has occurred. Please contact support.",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Login successful!",
		"success":    true,
		"session_id": sessionKey,
	})
}
