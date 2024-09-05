package handlers

import (
	"techinical/sentences"
	"techinical/users/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerDelete struct {
	sentences sentences.Sentences
}

func NewHandlerDeleteUser(sentences sentences.Sentences) HandlerDelete {
	if sentences == nil {
		panic("El repositorio de sentences es nil")
	}
	return HandlerDelete{sentences: sentences}
}

func (h *HandlerDelete) DeleteUser(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	user := dto.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return h.sentences.DeleteRecord(c, "users", "Id_User", user.Id_User, &user)
}
