// Package handlers ...
package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid session data"})
	}
	if err := h.repo.CreateSession(&sess); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create session"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "session created"})
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
	var sess models.Session
	if err := c.BodyParser(&sess); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid session data"})
	}
	valid, err := h.repo.VerifySession(sess.Get("username").(string), sess.Session.ID())
	if err != nil || !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid session"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "session valid"})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid session data"})
	}
	if err := h.repo.DeleteSession(sess.Get("username").(string), sess.Session.ID()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete session"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "session deleted"})
}
