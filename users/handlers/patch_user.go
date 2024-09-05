package handlers

import (
	"context"
	"strconv"
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
	user := dto.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if user.Id_User == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Debe ingresar el id de un usuario"})
	}
	// Completamos los campos faltantes usando IA
	if err := h.completeUserFields(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.updateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *HandlerPatch) completeUserFields(user *dto.User) error {
	if user.Name_User == "" {
		name, err := h.ChatGptApi.ChatGptMessague(context.Background(), "generame un nombre corto alazar")
		if err != nil {
			return err
		}
		user.Name_User = "Modificado: " + name
	}
	if user.Cell_Phone == 0 {
		output, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Genera unicamente un numero al azar de 10 dígitos")
		if err != nil {
			return err
		}

		cellPhone, err := strconv.Atoi(output)
		if err != nil {
			return err
		}
		user.Cell_Phone = cellPhone
	}
	return nil
}

func (h *HandlerPatch) updateUser(user *dto.User) (err error) {
	dbCon := db.ConectionDb()
	dbCon = dbCon.Exec(`UPDATE users SET name_user = ?, cell_phone = ? WHERE id_user = ?`, user.Name_User, user.Cell_Phone, user.Id_User)
	if dbCon.Error != nil {
		return dbCon.Error
	}
	return
}
