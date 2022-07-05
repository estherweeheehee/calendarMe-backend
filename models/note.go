package models

import ("gorm.io/gorm"
		"time"
)

type Note struct {
	ID        uint           `json:"id" gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`

	Month string `json:"month"`
	Date string 	`json:"date"`
	Title string   `json:"title"`
	Post  string   `json:"post"`
	Tag   string `json:"tag"`
}
