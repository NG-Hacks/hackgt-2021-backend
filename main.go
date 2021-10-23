package main
import (
	"https://github.com/ngngardner/hackgt-2021-backend/database"
)

func main() {
	db, err := database.connectDB()
	if err != nil {
		panic(err)
	}
}
