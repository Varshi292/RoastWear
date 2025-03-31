package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetImagesHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")
		if username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username is required.",
			})
		}

		var images []string
		err := db.Table("user_uploads").
			Where("username = ?", username).
			Pluck("filepath", &images).Error

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve image data.",
			})
		}

		if len(images) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No images found for this user.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"images": images,
		})
	}
}
