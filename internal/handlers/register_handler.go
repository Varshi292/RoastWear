package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	service *services.UserService
}

func NewRegisterHandler(service *services.UserService) *RegisterHandler {
	return &RegisterHandler{service: service}
}

func (handler *RegisterHandler) UserRegister(c *fiber.Ctx) error {
	var request models.UserRegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := handler.service.RegisterUser(&request); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user registered successfully"})
}
