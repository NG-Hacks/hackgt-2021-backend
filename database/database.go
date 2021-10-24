package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	latitude  int64
	longitude int64
	altitude  int64
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

func AddGeoLocation(db *gorm.DB, _latitude int64, _longitude int64, _altitude int64) *gorm.DB {

	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// result := db.Create(&user) // pass pointer of data to Create
	result := db.Select("latitude", "longitude", "altitude").Create(&User{latitude: _latitude, longitude: _longitude, altitude: _altitude})
	return result
}
