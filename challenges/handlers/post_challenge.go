package handlers

import (
	"context"
	"errors"
	"strconv"
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
	challenge := dto.Challenge{}
	if err := c.BodyParser(&challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.validateAndFillChallengeFields(&challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.saveChallenge(&challenge); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(challenge)
}

func (h *HandlerPost) validateAndFillChallengeFields(challenge *dto.Challenge) error {
	if challenge.Id_Video == "" {
		return errors.New("se requiere ingresar el id de un video")
	}
	if challenge.Id_Challenge == "" {
		challenge.Id_Challenge = uuid.NewString()
	}

	if challenge.Number_Participants == 0 {
		output, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Genera unicamente un numero entre 1 y 100")
		if err != nil {
			return err
		}
		num, err := strconv.Atoi(output)
		if err != nil {
			return err
		}
		challenge.Number_Participants = num
	}
	if challenge.Name_Challenge == "" {
		name, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Generame un nombre que no pase de 50 caracteres")
		if err != nil {
			return err
		}
		challenge.Name_Challenge = name
	}
	return nil
}

func (h *HandlerPost) saveChallenge(challenge *dto.Challenge) (err error) {
	dbCon := db.ConectionDb()
	dbCon = dbCon.Exec(
		`INSERT INTO challenges(id_challenge, id_video, name_challenge, number_participants) VALUES (?, ?, ?, ?)`,
		challenge.Id_Challenge, challenge.Id_Video, challenge.Name_Challenge, challenge.Number_Participants,
	)
	if dbCon.Error != nil {
		return dbCon.Error
	}
	return
}
