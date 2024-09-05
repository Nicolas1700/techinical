package handlers

import (
	"techinical/sentences"
	"techinical/video/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerDelete struct {
	sentences sentences.Sentences
}

func NewHandlerDeleteVideo(sentences sentences.Sentences) HandlerDelete {
	if sentences == nil {
		panic("El repositorio de sentences es nil")
	}
	return HandlerDelete{sentences: sentences}
}

func (h *HandlerDelete) DeleteVideo(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	video := dto.Video{}
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return h.sentences.DeleteRecord(c, "video", "id_video", video.Id_Video, &video)
}
