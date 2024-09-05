package handlers

import (
	"context"
	"errors"
	"techinical/db"
	sharedRepo "techinical/shared/repository"
	"techinical/video/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerPatch struct {
	ChatGptApi sharedRepo.ChatGptApi
}

func NewHandlerPatchVideo(chatGptApi sharedRepo.ChatGptApi) HandlerPatch {
	if chatGptApi == nil {
		panic("El repositorio es nil")
	}
	return HandlerPatch{ChatGptApi: chatGptApi}
}

func (h *HandlerPatch) PatchVideo(c *fiber.Ctx) error {
	video := dto.Video{}
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.validateAndFillVideoFields(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.updateVideo(&video); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(video)
}

func (h *HandlerPatch) validateAndFillVideoFields(video *dto.Video) error {
	if video.Id_Video == "" {
		return errors.New("es necesario recibir el id del video")
	}

	if video.Name_Video == "" {
		name, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Generame un nombre corto al azar")
		if err != nil {
			return err
		}
		video.Name_Video = "Modificado: " + name
	}
	if video.Url_Video == "" {
		url, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Generame solo una url de ejemplo corta")
		if err != nil {
			return err
		}
		video.Url_Video = url
	}
	return nil
}

func (h *HandlerPatch) updateVideo(video *dto.Video) (err error) {
	dbCon := db.ConectionDb()
	dbCon = dbCon.Exec(`UPDATE video SET id_user = ?, name_video = ?, url_video = ? WHERE id_video = ?`,
		video.Id_User, video.Name_Video, video.Url_Video, video.Id_Video)

	if dbCon.Error != nil {
		return dbCon.Error
	}
	return
}
