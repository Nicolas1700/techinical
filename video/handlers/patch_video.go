package handlers

import (
	"context"
	"fmt"
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
	// Obtenemos los datos del body
	video := dto.Video{}
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	result, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Generame un chiste")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	fmt.Println("Chatgpt", result)
	// Asiganos un uuid si no llega un valor
	if video.Id_Video == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Es necesario recibir el id para actualizar al usuario",
		})
	}
	dbCon := db.ConectionDb()
	// Definimos la consulta
	dbCon.Exec(`UPDATE video SET id_user = ?, name_video = ?,  url_video = ? WHERE id_video = ?`, video.Id_User, video.Name_Video, video.Url_Video, video.Id_Video)
	if dbCon.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": dbCon.Error,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(video)
}
