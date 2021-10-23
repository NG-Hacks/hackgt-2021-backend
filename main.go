package main

func main() {
	db, err := database.connectDB()
	if err != nil {
		panic(err)
	}
}
