package handlers

import (
	"techinical/sentencias"
	"techinical/video/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerDelete struct {
}

func NewHandlerDeleteVideo() HandlerDelete {
	return HandlerDelete{}
}

func (h *HandlerDelete) DeleteVideo(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	video := dto.Video{}
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return sentencias.DeleteRecord(c, "video", "id_video", video.Id_Video, &video)
}
