package main

import (
	"fmt"

	"github.com/ngngardner/hackgt-2021-backend/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Cannot connect to db")
	}
	db.AutoMigrate(&database.User{})

	//latitude := pgtype.Float8.Float(22.1)
	result := database.AddGeoLocation(db, 22.1, 10.2, 100.1)
	fmt.Println(result)
}
