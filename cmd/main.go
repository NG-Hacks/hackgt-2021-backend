package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/ngngardner/hackgt-2021-backend/database"
)

func main() {
	app := fiber.New()
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Cannot connect to db")
	}
	db.AutoMigrate(&database.User{})

	app.Listen(":8080")
}
