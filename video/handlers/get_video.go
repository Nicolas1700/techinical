package handlers

import (
	"techinical/db"
	"techinical/sentences"
	"techinical/video/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerGet struct {
	sentences sentences.Sentences
}

func NewHandlerGetVideo(sentences sentences.Sentences) HandlerGet {
	if sentences == nil {
		panic("El repositorio de sentences es nil")
	}
	return HandlerGet{sentences: sentences}
}

func (h *HandlerGet) GetVideo(c *fiber.Ctx) error {
	query, err := h.sentences.PaginateAndQuery(c, "video")
	if err != nil {
		return err
	}
	dbCon := db.ConectionDb()
	data := []dto.Video{}
	dbCon.Raw(query).Scan(&data)
	if len(data) > 0 {
		return c.Status(200).JSON(data)
	}

	return fiber.NewError(fiber.StatusNoContent, "No content")
}
