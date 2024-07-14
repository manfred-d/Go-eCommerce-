package database

import (
	"backend/go_backend/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) (*gorm.DB, error) {

	// var err error
	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})

	if err != nil {
		log.Fatalf("error %v", err)
	}

	return db, nil
}

// func MigrateDB(db *gorm.DB) error {
// 	err := db.AutoMigrate()

// 	if err != nil {
// 		log.Fatalf("Error %v", err)
// 	}

// 	return nil
// }
