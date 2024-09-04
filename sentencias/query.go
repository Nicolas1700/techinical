package sentencias

import (
	"fmt"
	"strconv"
	"techinical/db"

	"github.com/gofiber/fiber/v2"
)

func PaginateAndQuery(c *fiber.Ctx, tableName string) (string, error) {
	// Valor por defecto '1'
	page := c.Query("page", "1")
	// Valor por defecto '10'
	limit := c.Query("limit", "10")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return "", c.Status(400).JSON(fiber.Map{
			"error": "Invalid page parameter",
		})
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return "", c.Status(400).JSON(fiber.Map{
			"error": "Invalid limit parameter",
		})
	}
	offset := strconv.Itoa((pageInt - 1) * limitInt)
	// Definimos la consulta
	return `SELECT * FROM ` + tableName + ` LIMIT ` + limit + ` OFFSET ` + offset, nil
}

func DeleteRecord(c *fiber.Ctx, tableName string, idColumn string, idValue string, dest interface{}) error {
	if idValue == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Es necesario recibir el id para eliminar el registro",
		})
	}

	dbCon := db.ConectionDb()
	// Definimos la consulta
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableName, idColumn)
	dbCon = dbCon.Exec(query, idValue)

	if dbCon.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": dbCon.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dest)
}
