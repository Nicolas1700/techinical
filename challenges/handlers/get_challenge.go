package handlers

import (
	"techinical/challenges/infrastructura/dto"
	"techinical/db"
	"techinical/sentencias"

	"github.com/gofiber/fiber/v2"
)

type HandlerGet struct {
}

func NewHandlerGetChallenge() HandlerGet {
	return HandlerGet{}
}

func (h *HandlerGet) GetChallenge(c *fiber.Ctx) error {
	query, err := sentencias.PaginateAndQuery(c, "challenges")
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
