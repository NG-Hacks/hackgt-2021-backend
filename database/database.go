package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
	IsSharing bool    `json:"isSharing"`
}

func ConnectDB() (*gorm.DB, error) {
	// dsn := "host=35.192.7.243 port=5432 user=postgres dbname=compass password=abc123 sslmode=disable"
	dsn := "host=localhost port=5432 user=postgres dbname=postgres password=abc123 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateUser(db *gorm.DB, user User) *gorm.DB {
	// Create user
	result := db.Create(&user)
	return result
}

func ReadUser(db *gorm.DB, id int) User {
	// Read user from db
	var user User
	db.Where("id = ?", id).Find(user)
	return user
}

func SetIsSharing(db *gorm.DB, user User, _isSharing bool) {
	// choose if user wants to share location
	db.Model(&User{}).Where("id = ?", user.ID).Updates(User{IsSharing: _isSharing})
}

func UpdateGPScoordinates(db *gorm.DB, _latitude float32, _longitude float32, user User) {
	db.Model(&User{}).Where("id = ?", user.ID).Updates(User{Latitude: _latitude, Longitude: _longitude})

}

func UpdateAltitude(db *gorm.DB, _altitude float32, _longitude float32, user User) {
	db.Model(&User{}).Where("id = ?", user.ID).Updates(User{Altitude: _altitude})
}

func UpdateUser(db *gorm.DB, user User) {
	// update user
	db.Save(&user)
}
