package database

import (
	"log"
	"smedia/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dbURL := "postgres://pg:pass@localhost:7070/crud"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	return db, nil

}
