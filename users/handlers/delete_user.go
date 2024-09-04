package handlers

import (
	"techinical/sentencias"
	"techinical/users/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerDelete struct {
}

func NewHandlerDeleteUser() HandlerDelete {
	return HandlerDelete{}
}

func (h *HandlerDelete) DeleteUser(c *fiber.Ctx) error {
	// Obtenemos los datos del body
	user := dto.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return sentencias.DeleteRecord(c, "users", "Id_User", user.Id_User, &user)
}
