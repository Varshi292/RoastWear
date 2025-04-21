package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type RegisterHandler struct {
	userService *services.UserService
	sessionRepo *repositories.SessionRepository
}

func NewRegisterHandler(userService *services.UserService, sessionRepo *repositories.SessionRepository) *RegisterHandler {
	return &RegisterHandler{
		userService: userService,
		sessionRepo: sessionRepo,
	}
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
			"message": "Invalid request format!",
			"details": err.Error(),
		})
	}
	err := validate.Struct(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			field := err.Field()
			tag := err.Tag()
			switch field {
			case "Username":
				switch tag {
				case "required":
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"message": "Username is required!",
						"details": err,
					})
				case "min", "max":
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"message": "Username must be between 3 and 20 characters.",
						"details": err,
					})
				}
			case "Email":
				switch tag {
				case "required":
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"message": "Email is required!",
						"details": err,
					})
				case "email":
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"message": "Invalid email address format.",
						"details": err,
					})
				}
			case "Password":
				switch tag {
				case "required":
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"message": "Password is required!",
						"details": err,
					})
				case "min", "max":
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"message": "Password must be between 8 and 128 characters.",
						"details": err,
					})
				}
			}
		}
	}

	if err := handler.userService.RegisterUser(&request); err != nil {
		if utils.NewErrUserExists(request.Username).Error() == err.Error() {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Username '" + request.Username + "' is already taken.",
				"details": err.Error(),
			})
		}
		if utils.NewErrEmailExists(request.Email).Error() == err.Error() {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Email address '" + request.Email + "' is registered to an account.",
				"details": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error has occurred. Please contact support.",
			"details": err.Error(),
		})
	}

	sessionKey, err := utils.StartSession(c, handler.sessionRepo, request.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error has occurred. Please contact support.",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "User registered successfully!",
		"success":    true,
		"session_id": sessionKey,
	})
}
