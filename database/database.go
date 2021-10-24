package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Latitude  float32    `json:"latitude"`
	Longitude float32    `json:"longitude"`
	Altitude  float32    `json:"altitude"`
	IsSharing bool       `json:"isSharing"`
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=35.192.7.243 port=5432 user=postgres dbname=compass password=abc123 sslmode=disable"
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

func ReadUser(db *gorm.DB, id int) *gorm.DB {
	// Read user from db
	return db.First(&User{}, "id = ?", id)
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
