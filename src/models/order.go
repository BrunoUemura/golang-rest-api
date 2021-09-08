package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          int   `json:"id" gorm:"primaryKey"`
	UserId      int   `json:"userId"` 
	User        User  `gorm:"foreignKey:UserId"`
	BookId      int   `json:"bookId"` 
	Book		Book  `gorm:"foreignKey:BookId"`
	Quantity    int    `json:"quantity"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}