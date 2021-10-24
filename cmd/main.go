package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GeoLocation struct {
	gorm.Model
	userId    int
	lattitude float32
	longitude float32
	altitude  float32
}

func connectDB() (*gorm.DB, error) {
	// host=hackgt-329917:us-central1:compass port=5432 user=postgres dbname=compass password=abc123 sslmode=disable
	//db, err := gorm.Open("postgres", "host=hackgt-329917:us-central1:compass port=5432 user=postgres dbname=compass password=abc123 sslmode=disable")

	dsn := "host=35.192.7.243 port=5432 user=postgres dbname=compass password=abc123 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func addGeoLocation(db *gorm.DB, _userId int, _lattitude float32, _longitude float32, _altitude float32) {
	db.Create(&GeoLocation{userId: _userId, lattitude: _lattitude, longitude: _longitude, altitude: _altitude})
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Cannot connect to db")
	}

	addGeoLocation(db, 1, 23.42, 45.62, 100.0)
}
