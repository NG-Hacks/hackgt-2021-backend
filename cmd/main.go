package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ngngardner/hackgt-2021-backend/database"
	"github.com/ngngardner/hackgt-2021-backend/handlers"
)

func main() {
	// db
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Cannot connect to db")
	}
	db.AutoMigrate(&database.User{})

	// fiber
	app := fiber.New()
	app.Get("/user/:id", handlers.GetUser)
	app.Listen(":8080")
}
