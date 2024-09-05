package handlers

import (
	"context"
	"strconv"
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
	user := dto.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if user.Id_User == "" {
		user.Id_User = uuid.NewString()
	}

	if err := h.populateUserFields(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.saveUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *HandlerPost) populateUserFields(user *dto.User) (err error) {
	if user.Cell_Phone == 0 {
		if err = h.addCellPhoneWithIa(user); err != nil {
			return
		}
	}
	if user.Name_User == "" {
		if err = h.addNameWithIa(user); err != nil {
			return
		}
	}
	return
}

func (h *HandlerPost) addCellPhoneWithIa(user *dto.User) (err error) {
	output, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Genera unicamente un numero alazar de 10 digitos")
	if err != nil {
		return
	}
	cellPhone, err := strconv.Atoi(output)
	if err != nil {
		return
	}
	user.Cell_Phone = cellPhone
	return nil
}

func (h *HandlerPost) addNameWithIa(user *dto.User) (err error) {
	name, err := h.ChatGptApi.ChatGptMessague(context.Background(), "Generame un nombre que no pase de 50 caracteres")
	if err != nil {
		return
	}
	user.Name_User = name
	return nil
}

func (h *HandlerPost) saveUser(user *dto.User) (err error) {
	dbCon := db.ConectionDb()
	dbCon = dbCon.Exec(`INSERT INTO users(id_user, name_user, cell_phone) VALUES (?, ?, ?)`, user.Id_User, user.Name_User, user.Cell_Phone)
	if dbCon.Error != nil {
		return dbCon.Error
	}
	return
}
