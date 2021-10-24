package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ngngardner/hackgt-2021-backend/database"
)

func GetUser(c *fiber.Ctx) error {
	// get user from db
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	idStr, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	user := database.ReadUser(db, idStr)
	return c.JSON(user)
}
