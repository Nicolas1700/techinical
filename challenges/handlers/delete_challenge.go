package handlers

import (
	"techinical/challenges/infrastructura/dto"
	"techinical/sentences"

	"github.com/gofiber/fiber/v2"
)

type HandlerDelete struct {
	sentences sentences.Sentences
}

func NewHandlerDeleteChallenge(sentences sentences.Sentences) HandlerDelete {
	if sentences == nil {
		panic("El repositorio de sentences es nil")
	}
	return HandlerDelete{sentences: sentences}
}

func (h *HandlerDelete) DeleteChallenge(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	challenge := dto.Challenge{}
	if err := c.BodyParser(&challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return h.sentences.DeleteRecord(c, "challenges", "id_challenge", challenge.Id_Challenge, &challenge)
}
