package config

import (
	"calendarme-backend/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
  var DB *gorm.DB

  func Connect() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&models.Note{})
	DB = db
  }
