// Package unitTest provides integration and unit tests for RoastWear backend services.
package unitTest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repository"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/session"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func setupLoginTestApp() (*fiber.App, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&models.User{}, &models.Session{})

	session.InitializeSessionStore(config.SessionConfig{
		Key:         "cookie:session_id",
		Path:        "/",
		MaxAge:      3600 * 1e9,
		HttpOnly:    true,
		Secure:      false,
		SameSite:    "Lax",
		SessionOnly: false,
	})

	userRepo := &repository.UserRepository{Db: db}
	sessionService := services.NewSessionService(db)
	authService := services.NewAuthService(userRepo, sessionService)

	app := fiber.New()
	app.Post("/login", handlers.NewLoginHandler(authService, sessionService).UserLogin)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	db.Create(&models.User{
		Username: "validUser",
		Password: string(hashedPassword),
	})

	return app, db
}

func loginAndGetSessionID(t *testing.T, app *fiber.App) string {
	loginReq := models.UserLoginRequest{
		Username: "validUser",
		Password: "testpass",
	}
	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&result)
	sessionID := result["session_id"].(string)
	require.NotEmpty(t, sessionID)
	return sessionID
}

func TestSessionVerificationAfterLogin(t *testing.T) {
	app, db := setupLoginTestApp()
	sessionID := loginAndGetSessionID(t, app)

	verifyApp := fiber.New()
	verifyApp.Post("/session/verify", handlers.NewSessionHandler(services.NewSessionService(db)).VerifySession)
	verifyBody, _ := json.Marshal(models.Session{
		Username:  "validUser",
		SessionID: sessionID,
	})
	verifyReq := httptest.NewRequest("POST", "/session/verify", bytes.NewReader(verifyBody))
	verifyReq.Header.Set("Content-Type", "application/json")
	verifyResp, err := verifyApp.Test(verifyReq)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, verifyResp.StatusCode)
}

func TestSessionStillValidAfterUserDeletion(t *testing.T) {
	app, db := setupLoginTestApp()
	sessionID := loginAndGetSessionID(t, app)

	db.Delete(&models.User{}, "username = ?", "validUser")

	verifyApp := fiber.New()
	verifyApp.Post("/session/verify", handlers.NewSessionHandler(services.NewSessionService(db)).VerifySession)
	verifyBody, _ := json.Marshal(models.Session{
		Username:  "validUser",
		SessionID: sessionID,
	})
	verifyReq := httptest.NewRequest("POST", "/session/verify", bytes.NewReader(verifyBody))
	verifyReq.Header.Set("Content-Type", "application/json")
	verifyResp, err := verifyApp.Test(verifyReq)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, verifyResp.StatusCode)
}

func TestTamperedSessionIDFailsVerification(t *testing.T) {
	app, db := setupLoginTestApp()
	sessionID := loginAndGetSessionID(t, app)

	verifyApp := fiber.New()
	verifyApp.Post("/session/verify", handlers.NewSessionHandler(services.NewSessionService(db)).VerifySession)
	tamperedBody, _ := json.Marshal(models.Session{
		Username:  "validUser",
		SessionID: sessionID + "_hacked",
	})
	tamperedReq := httptest.NewRequest("POST", "/session/verify", bytes.NewReader(tamperedBody))
	tamperedReq.Header.Set("Content-Type", "application/json")
	tamperedResp, err := verifyApp.Test(tamperedReq)

	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, tamperedResp.StatusCode)
}
