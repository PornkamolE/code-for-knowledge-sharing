package handlers

import (
	"github.com/gofiber/fiber/v2"
	"my-knowledge-sharing/app/core/services"
	"strconv"
)

type UserHandler struct {
	svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	id, err := h.svc.CreateUser(c.Context(), req.Name)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": id})
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	name, err := h.svc.GetUser(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"id":   id,
		"name": name,
	})
}