package database

import (
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/gorm"
)

type GeoLocation struct {
	gorm.Model
	userId    int
	lattitude float32
	longitude float32
	altitude  float32
}

func connectDB() (*gorm.DB, err){
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func addGeoLocation(db *gorm.DB,  _userId int,  _lattitude float32,  _longitude float32, _altitude float32 {
	db.Create(&GeoLocation{userId: _userId, lattitude: _lattitude, longitude: _longitude, altitude: _altitude})
}


