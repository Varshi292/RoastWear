// Package handlers ...
package handlers

import (
	"RoastWear/internal/models"
	"RoastWear/internal/services"
	"RoastWear/internal/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
)

// LoginHandler ...
//
// Fields:
//   - service: ...
type LoginHandler struct {
	service *services.AuthService
}

// NewLoginHandler ...
//
// Parameters:
//   - service: ...
//
// Returns:
//   - *LoginHandler: ...
func NewLoginHandler(service *services.AuthService) *LoginHandler {
	return &LoginHandler{service: service}
}

// UserLogin handles the login process for a user.
// @Summary User login
// @Description Authenticates a user with their credentials (username and password). Returns a session token on successful login. In case of errors, it provides appropriate status codes and error messages for invalid credentials, missing fields, or server issues.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserLoginRequest true "User login credentials (username and password)"
// @Success 200 {object} map[string]interface{} "Login successful, returns a success message"
// @Failure 400 {object} map[string]interface{} "Bad request, invalid request format"
// @Failure 422 {object} map[string]interface{} "Validation error, missing required fields (username and password)"
// @Failure 401 {object} map[string]interface{} "Unauthorized, invalid username or password"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /login [post]
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
	if err := handler.service.LoginUser(&request, c); err != nil {
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successful"})
}
