package handlers

import (
	"context"
	"fmt"
	"techinical/db"
	sharedRepo "techinical/shared/repository"
	"techinical/users/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HandlerPost struct {
	ChatGptApi sharedRepo.ChatGptApi
}

func NewHandlerPostUser(chatGptApi sharedRepo.ChatGptApi) HandlerPost {
	if chatGptApi == nil {
		panic("El repositorio es nil")
	}
	return HandlerPost{ChatGptApi: chatGptApi}
}

func (h *HandlerPost) PostUser(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	user := dto.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	// Asiganos un uuid si no llega un valor en este
	if user.Id_User == "" {
		user.Id_User = uuid.NewString()
	}
	dbCon := db.ConectionDb()
	resultado, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Estoy ablando con chatgpt???")
	if err != nil {
		return err
	}
	fmt.Println("resultado", resultado)
	// Definimos la consulta
	dbCon = dbCon.Exec(`INSERT INTO users(id_user, name_user, cell_phone) VALUES (?, ?, ?)`,
		user.Id_User, user.Name_User, user.Cell_Phone,
	)
	if dbCon.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": dbCon.Error,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
