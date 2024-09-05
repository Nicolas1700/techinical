package handlers

import (
	"techinical/challenges/infrastructura/dto"
	"techinical/db"
	"techinical/sentences"

	"github.com/gofiber/fiber/v2"
)

type HandlerGet struct {
	sentences sentences.Sentences
}

func NewHandlerGetChallenge(sentences sentences.Sentences) HandlerGet {
	if sentences == nil {
		panic("El repositorio de sentences es nil")
	}
	return HandlerGet{sentences: sentences}
}

func (h *HandlerGet) GetChallenge(c *fiber.Ctx) error {
	query, err := h.sentences.PaginateAndQuery(c, "challenges")
	if err != nil {
		return err
	}
	dbCon := db.ConectionDb()
	challenges := []dto.Challenge{}
	dbCon.Raw(query).Scan(&challenges)
	if len(challenges) > 0 {
		return c.Status(200).JSON(challenges)
	}
	return fiber.NewError(fiber.StatusNoContent, "No content")
}
