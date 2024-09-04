package handlers

import (
	"techinical/db"
	"techinical/sentencias"
	"techinical/users/infrastructura/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerGet struct {
}

func NewHandlerGetUser() HandlerGet {
	return HandlerGet{}
}

func (h *HandlerGet) GetUser(c *fiber.Ctx) error {
	query, err := sentencias.PaginateAndQuery(c, "users")
	if err != nil {
		return err
	}
	dbCon := db.ConectionDb()
	users := []dto.User{}
	dbCon.Raw(query).Scan(&users)
	if len(users) > 0 {
		return c.Status(200).JSON(users)
	}
	return fiber.NewError(fiber.StatusNoContent, "No content")
}
