package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          int   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MediumPrice float64 `json:"medium_price"`
	Author      string `json:"author"`
	ImageURL    string `json:"img_url"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}