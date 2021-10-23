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

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN: "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	  })

}

// func addGeoLocation(db *gorm.DB,  _userId int,  _lattitude float32,  _longitude float32, _altitude float32 {
// 	db.Create(&GeoLocation{userId: _userId, lattitude: _lattitude, longitude: _longitude, altitude: _altitude})
// }
