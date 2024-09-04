package handlers

import (
	"techinical/db"
	"techinical/sentencias"
	"techinical/video/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerGet struct {
}

func NewHandlerGetVideo() HandlerGet {
	return HandlerGet{}
}

func (h *HandlerGet) GetVideo(c *fiber.Ctx) error {
	query, err := sentencias.PaginateAndQuery(c, "video")
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
