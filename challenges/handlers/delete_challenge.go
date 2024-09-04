package handlers

import (
	"techinical/challenges/infrastructura/dto"
	"techinical/sentencias"

	"github.com/gofiber/fiber/v2"
)

type HandlerDelete struct {
}

func NewHandlerDeleteChallenge() HandlerDelete {
	return HandlerDelete{}
}

func (h *HandlerDelete) DeleteChallenge(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	challenge := dto.Challenge{}
	if err := c.BodyParser(&challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return sentencias.DeleteRecord(c, "challenges", "id_challenge", challenge.Id_Challenge, &challenge)
}
