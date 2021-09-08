package migrations

import (
	"github.com/BrunoUemura/golang-rest-api/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.Order{})
}