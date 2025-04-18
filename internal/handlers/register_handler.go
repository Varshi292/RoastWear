// Package handlers ...
package handlers

import (
	"RoastWear/internal/models"
	"RoastWear/internal/services"
	"RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// RegisterHandler ...
//
// Fields:
//   - service: ...
type RegisterHandler struct {
	service *services.UserService
}

// NewRegisterHandler ...
//
// Parameters:
//   - service: ...
//
// Returns:
//   - *RegisterHandler: ...
func NewRegisterHandler(service *services.UserService) *RegisterHandler {
	return &RegisterHandler{service: service}
}

// UserRegister handles the user registration process.
// @Summary User registration
// @Description Registers a new user by creating an account with the provided username, email, and password. If the user already exists (either by username or email), an error is returned. If there are missing fields, a validation error is returned. Includes a detailed response for successful registration or error scenarios.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserRegisterRequest true "User registration details (username, email, and password)"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Bad request, invalid request format"
// @Failure 422 {object} map[string]interface{} "Validation error, missing required fields (username, email, or password)"
// @Failure 409 {object} map[string]interface{} "Conflict error, username or email already exists"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /register [post]
func (handler *RegisterHandler) UserRegister(c *fiber.Ctx) error {
	var request models.UserRegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid request format",
			"details": err.Error(),
		})
	}
	if request.Username == "" || request.Email == "" || request.Password == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "validation error",
			"details": "username, email, and password are required",
		})
	}
	if err := handler.service.RegisterUser(&request); err != nil {
		if utils.NewErrUserExists(request.Username).Error() == err.Error() || utils.NewErrEmailExists(request.Email).
			Error() == err.Error() {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error":   "conflict error",
				"details": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "internal server error",
			"details": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user registered successfully",
		"user":    request.Username,
	})
}
