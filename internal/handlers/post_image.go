package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"time"
)

var ImageDir = "./user_images" // Adjust path as needed

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

		if file.Header.Get("Content-Type") != "image/png" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Only PNG images are allowed.",
			})
		}

		// Ensure the directory exists
		if _, err := os.Stat(ImageDir); os.IsNotExist(err) {
			os.MkdirAll(ImageDir, 0755)
		}

		uniqueName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
		savePath := filepath.Join(ImageDir, uniqueName)

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(savePath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		// Save to DB using GORM
		err = db.Exec("INSERT INTO user_uploads (username, filepath) VALUES (?, ?)", username, uniqueName).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to store image data.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "Image uploaded and data saved.",
			"filepath": uniqueName,
		})
	}
}
