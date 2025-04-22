package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetImagesHandler retrieves images uploaded by a user.
// @Summary Get uploaded images for a user
// @Description This endpoint allows retrieving a list of filepaths for the images uploaded by a user, based on the username.
// @Tags Images
// @Accept json
// @Produce json
// @Param username query string true "Username of the user"
// @Success 200 {object} map[string][]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /images [get]

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
