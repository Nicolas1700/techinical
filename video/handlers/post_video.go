package handlers

import (
	"context"
	"errors"
	"techinical/db"
	sharedRepo "techinical/shared/repository"
	"techinical/video/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HandlerPost struct {
	ChatGptApi sharedRepo.ChatGptApi
}

func NewHandlerPostVideo(chatGptApi sharedRepo.ChatGptApi) HandlerPost {
	if chatGptApi == nil {
		panic("El repositorio es nil")
	}
	return HandlerPost{ChatGptApi: chatGptApi}
}

func (h *HandlerPost) PostVideo(c *fiber.Ctx) error {
	video := dto.Video{}
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.validateVideoFields(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.saveVideo(&video); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(video)
}

func (h *HandlerPost) validateVideoFields(video *dto.Video) error {
	if video.Id_Video == "" {
		video.Id_Video = uuid.NewString()
	}
	if video.Id_User == "" {
		return errors.New("debe ingresar un id de usuario")
	}

	if video.Name_Video == "" {
		name, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Generame solo un nombre que no pase de 50 caracteres")
		if err != nil {
			return err
		}
		video.Name_Video = name
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

func (h *HandlerPost) saveVideo(video *dto.Video) (err error) {
	dbCon := db.ConectionDb()
	dbCon = dbCon.Exec(`INSERT INTO video(id_video, id_user, name_video, url_video) VALUES (?, ?, ?, ?)`,
		video.Id_Video, video.Id_User, video.Name_Video, video.Url_Video)

	if dbCon.Error != nil {
		return dbCon.Error
	}
	return
}
