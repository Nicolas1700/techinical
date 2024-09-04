package handlers

import (
	"techinical/challenges/infrastructura/dto"
	"techinical/db"
	sharedRepo "techinical/shared/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HandlerPost struct {
	ChatGptApi sharedRepo.ChatGptApi
}

func NewHandlerPostChallenge(chatGptApi sharedRepo.ChatGptApi) HandlerPost {
	if chatGptApi == nil {
		panic("El repositorio es nil")
	}
	return HandlerPost{ChatGptApi: chatGptApi}
}

func (h *HandlerPost) PostChallenge(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	challenge := dto.Challenge{}
	if err := c.BodyParser(&challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	// Asiganos un uuid si no llega un valor en este
	if challenge.Id_Challenge == "" {
		challenge.Id_Challenge = uuid.NewString()
	}
	dbCon := db.ConectionDb()
	// Definimos la consulta
	dbCon = dbCon.Exec(`INSERT INTO challenges(id_challenge, id_video, name_challenge, number_participants) VALUES (?, ?, ?, ?)`,
		challenge.Id_Challenge, challenge.Id_Video, challenge.Name_Challenge, challenge.Number_Participants,
	)
	if dbCon.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": dbCon.Error,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(challenge)
}
