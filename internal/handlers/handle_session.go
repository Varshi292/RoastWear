// Package handlers ...
package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

// SessionHandler ...
//
// Fields:
//   - service: SessionRepository for managing session records
type SessionHandler struct {
	service *repositories.SessionRepository
}

// NewSessionHandler ...
//
// Parameters:
//   - service: pointer to SessionRepository
//
// Returns:
//   - *SessionHandler: a new instance of SessionHandler
func NewSessionHandler(service *repositories.SessionRepository) *SessionHandler {
	return &SessionHandler{service: service}
}

// CreateSession ...
// @Summary Create a session
// @Description Stores a new session for a user
// @Accept json
// @Produce json
// @Param session body models.Session true "Session details (username, sessionID)"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /session/create [post]
func (h *SessionHandler) CreateSession(c *fiber.Ctx) error {
	var session models.Session
	if err := c.BodyParser(&session); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid session data"})
	}
	if err := h.service.CreateSession(&session); err != nil {
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
	var session models.Session
	if err := c.BodyParser(&session); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid session data"})
	}
	valid, err := h.service.VerifySession(session.Username, session.SessionID)
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
	var session models.Session
	if err := c.BodyParser(&session); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid session data"})
	}
	if err := h.service.DeleteSession(session.Username, session.SessionID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete session"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "session deleted"})
}
