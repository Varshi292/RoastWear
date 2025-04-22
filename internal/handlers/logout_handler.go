package handlers

import (
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/gofiber/fiber/v2"
)

// LogoutHandler ...
type LogoutHandler struct {
	sessionRepo *repositories.SessionRepository
}

// NewLogoutHandler ...
func NewLogoutHandler(sessionRepo *repositories.SessionRepository) *LogoutHandler {
	return &LogoutHandler{
		sessionRepo: sessionRepo,
	}
}

// UserLogout handles the login process for a user.
// @Summary      Logs out a user
// @Description  Destroys the user session and logs the user out
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]interface{} "Logged out successfully"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /logout [post]
func (h *LogoutHandler) UserLogout(c *fiber.Ctx) error {
	sess, err := sessions.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Session not found.",
		})
	}
	if err := h.sessionRepo.DeleteSession(sess.ID()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete session.",
			"details": err.Error(),
		})
	}
	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to destroy session.",
			"details": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out successfully!",
		"success": true,
	})
}
