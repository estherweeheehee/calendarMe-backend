package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"calendarme-backend/models"
	
  )
  
  var DB *gorm.DB

  func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://estherweeheehee:password@localhost:5432/project4"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&models.Note{})
	DB = db
  }
