package handlers

import (
	"fmt"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var ImageDir = "./user_images"

func UploadImageHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		if username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username is required.",
			})
		}

		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No file uploaded.",
			})
		}

		// Check file extension instead of MIME (more consistent for tests too)
		if !strings.HasSuffix(strings.ToLower(file.Filename), ".png") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Only PNG images are allowed.",
			})
		}

		// Ensure the image directory exists
		if _, err := os.Stat(ImageDir); os.IsNotExist(err) {
			os.MkdirAll(ImageDir, 0755)
		}

		// Generate unique file name
		uniqueName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
		savePath := filepath.Join(ImageDir, uniqueName)

		// Open source file
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Create destination file
		dst, err := os.Create(savePath)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy contents
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		// Insert record into DB
		record := models.UserUpload{
			Username: username,
			Filepath: uniqueName,
		}
		if err := db.Create(&record).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to store image metadata.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "Image uploaded and data saved.",
			"filepath": uniqueName,
		})
	}
}
