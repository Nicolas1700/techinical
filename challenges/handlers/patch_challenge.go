package handlers

import (
	"context"
	"fmt"
	"techinical/challenges/infrastructura/dto"
	"techinical/db"
	sharedRepo "techinical/shared/repository"

	"github.com/gofiber/fiber/v2"
)

type HandlerPatch struct {
	ChatGptApi sharedRepo.ChatGptApi
}

func NewHandlerPatchChallenge(chatGptApi sharedRepo.ChatGptApi) HandlerPatch {
	if chatGptApi == nil {
		panic("El repositorio es nil")
	}
	return HandlerPatch{ChatGptApi: chatGptApi}
}

func (h *HandlerPatch) PatchChallenge(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	challenge := dto.Challenge{}
	if err := c.BodyParser(&challenge); err != nil {
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
	if challenge.Id_Challenge == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Es necesario recibir el id para actualizar al usuario",
		})
	}
	dbCon := db.ConectionDb()
	// Definimos la consulta
	dbCon.Exec(`UPDATE challenges SET id_video = ?, name_challenge = ?,  number_participants = ? WHERE id_challenge = ?`, challenge.Id_Video, challenge.Name_Challenge, challenge.Number_Participants, challenge.Id_Challenge)
	if dbCon.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": dbCon.Error,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(challenge)
}
