package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Latitude  float32
	Longitude float32
	Altitude  float32
}

func ConnectDB() (*gorm.DB, error) {
	// host=hackgt-329917:us-central1:compass port=5432 user=postgres dbname=compass password=abc123 sslmode=disable
	//db, err := gorm.Open("postgres", "host=hackgt-329917:us-central1:compass port=5432 user=postgres dbname=compass password=abc123 sslmode=disable")
	dsn := "host=35.192.7.243 port=5432 user=postgres dbname=compass password=abc123 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func AddGeoLocation(db *gorm.DB, _latitude float32, _longitude float32, _altitude float32) *gorm.DB {
	result := db.Create(&User{Latitude: _latitude, Longitude: _longitude, Altitude: _altitude})
	return result
}
