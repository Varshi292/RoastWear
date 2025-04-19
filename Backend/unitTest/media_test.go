package unitTest

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// setupTestApp initializes the Fiber app and in-memory SQLite DB
func setupTestApp() (*fiber.App, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// ✅ MIGRATE the user_uploads table used by the image handler
	if err := db.AutoMigrate(&models.UserUpload{}); err != nil {
		panic("failed to migrate test db: " + err.Error())
	}

	app := fiber.New()
	app.Post("/upload_image", handlers.UploadImageHandler(db))
	app.Get("/get_user_media", handlers.GetImagesHandler(db))
	return app, db
}

// createTestFileUploadRequest prepares a multipart/form-data request
func createTestFileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.Close()

	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

// TestUploadImage ensures a valid image is accepted
func TestUploadImage(t *testing.T) {
	app, _ := setupTestApp()

	imagePath := "test_data/test_image.png"
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		t.Fatal("Test image file does not exist")
	}

	req, err := createTestFileUploadRequest("/upload_image", map[string]string{"username": "testUser"}, "image", imagePath)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestUploadInvalidFile ensures non-image files are rejected
func TestUploadInvalidFile(t *testing.T) {
	app, _ := setupTestApp()

	invalidPath := "test_data/invalid.txt"
	if _, err := os.Stat(invalidPath); os.IsNotExist(err) {
		t.Fatal("Invalid test file does not exist")
	}

	req, err := createTestFileUploadRequest("/upload_image", map[string]string{"username": "testUser"}, "image", invalidPath)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	body := getResponseBody(t, resp)
	assert.Contains(t, body, "Only PNG images are allowed.")
}

// TestGetUserMediaSuccess tests fetching media for an existing user
func TestGetUserMediaSuccess(t *testing.T) {
	app, _ := setupTestApp()

	imagePath := "test_data/test_image.png"
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		t.Fatal("Test image file does not exist")
	}

	req, err := createTestFileUploadRequest("/upload_image", map[string]string{"username": "testUser"}, "image", imagePath)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	getReq := httptest.NewRequest("GET", "/get_user_media?username=testUser", nil)
	getResp, err := app.Test(getReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getResp.StatusCode)
	assert.Contains(t, getResponseBody(t, getResp), "images")
}

// TestGetUserMediaNotFound checks for non-existent user
func TestGetUserMediaNotFound(t *testing.T) {
	app, _ := setupTestApp()

	req := httptest.NewRequest("GET", "/get_user_media?username=nonExistingUser", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Contains(t, getResponseBody(t, resp), "No images found")
}

// Helper to read response body text
func getResponseBody(t *testing.T, resp *http.Response) string {
	bodyBytes, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	return string(bodyBytes)
}
