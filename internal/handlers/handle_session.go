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
// @Param session body models.SessionDoc true "Session data"
// @Success 201 {object} models.GenericResponse "Session successfully created"
// @Failure 400 {object} models.GenericResponse "Invalid session data"
// @Failure 500 {object} models.GenericResponse "Failed to create session"
// @Router /session/create [post]
func (h *SessionHandler) CreateSession(c *fiber.Ctx) error {
	var sess models.Session
	if err := c.BodyParser(&sess); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GenericResponse{
			Message: "Invalid request format!",
			Success: false,
		})
	}
	if err := h.repo.CreateSession(&sess); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.GenericResponse{
			Message: "Failed to create session!",
			Success: false,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(models.GenericResponse{
		Message: "Session created successfully!",
		Success: true,
	})
}

// VerifySession godoc
// @Summary Verify session
// @Description Validates a session by username and session ID
// @Tags Session
// @Accept json
// @Produce json
// @Param session body models.SessionDoc true "Session details"
// @Success 200 {object} models.GenericResponse
// @Failure 401 {object} models.GenericResponse
// @Router /session/verify [post]
func (h *SessionHandler) VerifySession(c *fiber.Ctx) error {
	sess, err := sessions.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GenericResponse{
			Message: "Invalid request format!",
			Success: false,
		})
	}
	if err := h.repo.GetSession(sess.ID()); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.GenericResponse{
			Message: "Session not found!",
			Success: false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.GenericResponse{
		Message: "Session found!",
		Success: true,
	})
}

// DeleteSession godoc
// @Summary Delete session
// @Description Removes a session from the database
// @Tags Session
// @Accept json
// @Produce json
// @Param session body models.SessionDoc true "Session details"
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.GenericResponse
// @Router /session/delete [delete]
func (h *SessionHandler) DeleteSession(c *fiber.Ctx) error {
	var sess models.Session
	if err := c.BodyParser(&sess); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GenericResponse{
			Message: "Invalid request format!",
			Success: false,
		})
	}
	if err := h.repo.DeleteSession(sess.SessionKey); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.GenericResponse{
			Message: "Failed to delete session!",
			Success: false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.GenericResponse{
		Message: "Session successfully deleted!",
		Success: true,
	})
}
