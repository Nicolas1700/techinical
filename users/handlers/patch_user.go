package handlers

import (
	"context"
	"fmt"
	"techinical/db"
	sharedRepo "techinical/shared/repository"
	"techinical/users/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerPatch struct {
	ChatGptApi sharedRepo.ChatGptApi
}

func NewHandlerPatchUser(chatGptApi sharedRepo.ChatGptApi) HandlerPatch {
	if chatGptApi == nil {
		panic("El repositorio es nil")
	}
	return HandlerPatch{ChatGptApi: chatGptApi}
}

func (h *HandlerPatch) PatchUser(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	user := dto.User{}
	if err := c.BodyParser(&user); err != nil {
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
	if user.Id_User == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Es necesario recibir el id para actualizar al usuario",
		})
	}
	dbCon := db.ConectionDb()
	// Definimos la consulta
	dbCon.Exec(`UPDATE users SET name_user = ?, cell_phone = ? WHERE id_user = ?`, user.Name_User, user.Cell_Phone, user.Id_User)
	if dbCon.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": dbCon.Error,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
