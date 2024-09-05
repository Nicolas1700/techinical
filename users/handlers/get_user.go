package handlers

import (
	"techinical/db"
	"techinical/sentences"
	"techinical/users/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerGet struct {
	sentences sentences.Sentences
}

func NewHandlerGetUser(sentences sentences.Sentences) HandlerGet {
	if sentences == nil {
		panic("El repositorio de sentences es nil")
	}
	return HandlerGet{sentences: sentences}
}

func (h *HandlerGet) GetUser(c *fiber.Ctx) error {
	query, err := h.sentences.PaginateAndQuery(c, "users")
	if err != nil {
		return err
	}
	dbCon := db.ConectionDb()
	users := []dto.User{}
	dbCon.Raw(query).Scan(&users)
	if len(users) > 0 {
		return c.Status(200).JSON(users)
	}
	return fiber.NewError(fiber.StatusNoContent, "No content")
}
