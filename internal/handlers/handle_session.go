// Package handlers ...
package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/gofiber/fiber/v2"
)

type SessionHandler struct {
	repo *repositories.SessionRepository
}

func NewSessionHandler(repo *repositories.SessionRepository) *SessionHandler {
	return &SessionHandler{repo: repo}
}

// CreateSession godoc
// @Summary Create a new user session
// @Description Stores a session in the database (used for login/session tracking)
// @Tags Session
// @Accept json
// @Produce json
// @Param session body models.Session true "Session data"
// @Success 201 {object} fiber.Map "Session successfully created"
// @Failure 400 {object} fiber.Map "Invalid session data"
// @Failure 500 {object} fiber.Map "Failed to create session"
// @Router /session/create [post]

func (h *SessionHandler) CreateSession(c *fiber.Ctx) error {
	var sess models.Session
	if err := c.BodyParser(&sess); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request format!",
			"details": err,
		})
	}
	if err := h.repo.CreateSession(&sess); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create session!",
			"details": err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Session created successfully!",
		"success": true,
	})
}

// VerifySession ...
// @Summary Verify session
// @Description Validates a session by username and session ID
// @Accept json
// @Produce json
// @Param session body models.Session true "Session details"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /session/verify [post]
func (h *SessionHandler) VerifySession(c *fiber.Ctx) error {
	sess, err := sessions.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request format!",
			"details": err.Error()})
	}
	if err := h.repo.GetSession(sess.ID()); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Session not found!",
			"details": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Session found!",
		"success": true,
	})
}

// DeleteSession ...
// @Summary Delete session
// @Description Removes a session from the database
// @Accept json
// @Produce json
// @Param session body models.Session true "Session details"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /session/delete [delete]
func (h *SessionHandler) DeleteSession(c *fiber.Ctx) error {
	var sess models.Session
	if err := c.BodyParser(&sess); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request format!",
			"details": err,
		})
	}
	if err := h.repo.DeleteSession(sess.SessionKey); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete session!",
			"details": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Session successfully deleted!",
		"success": true,
	})
}
