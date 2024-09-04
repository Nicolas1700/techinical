package handlers

import (
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
	// Obtenemos los datos del body
	video := dto.Video{}
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	// Asiganos un uuid si no llega un valor en este
	if video.Id_Video == "" {
		video.Id_Video = uuid.NewString()
	}
	dbCon := db.ConectionDb()
	// Definimos la consulta
	dbCon = dbCon.Exec(`INSERT INTO video(id_video, id_user, name_video, url_video) VALUES (?, ?, ?, ?)`,
		video.Id_Video, video.Id_User, video.Name_Video, video.Url_Video,
	)
	if dbCon.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": dbCon.Error,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(video)
}
