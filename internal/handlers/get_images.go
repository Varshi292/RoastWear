package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
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

		var uploads []models.UserUpload
		err := db.Where("username = ?", username).Find(&uploads).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve image data.",
			})
		}

		if len(uploads) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No images found for this user.",
			})
		}

		// Extract only the filepaths
		var filepaths []string
		for _, upload := range uploads {
			filepaths = append(filepaths, upload.Filepath)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"images": filepaths,
		})
	}
}
