package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/gofiber/fiber/v2"
)

type LoginHandler struct {
	service *services.AuthService
}

func NewLoginHandler(service *services.AuthService) *LoginHandler {
	return &LoginHandler{service: service}
}

func (handler *LoginHandler) UserLogin(c *fiber.Ctx) error {
	var request models.UserLoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := handler.service.LoginUser(&request, c); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "login successful"})

}
