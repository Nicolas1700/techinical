package handlers

import (
	"context"
	"errors"
	"strconv"
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
	challenge := dto.Challenge{}
	if err := c.BodyParser(&challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.validateAndFillChallengeFields(&challenge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.updateChallenge(&challenge); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(challenge)
}

func (h *HandlerPatch) validateAndFillChallengeFields(challenge *dto.Challenge) error {
	if challenge.Id_Challenge == "" {
		return errors.New("es necesario recibir el ID del desafío para actualizarlo")
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

func (h *HandlerPatch) updateChallenge(challenge *dto.Challenge) (err error) {
	dbCon := db.ConectionDb()
	dbCon = dbCon.Exec(
		`UPDATE challenges SET id_video = ?, name_challenge = ?, number_participants = ? WHERE id_challenge = ?`,
		challenge.Id_Video, challenge.Name_Challenge, challenge.Number_Participants, challenge.Id_Challenge,
	)
	if dbCon.Error != nil {
		return dbCon.Error
	}
	return
}
