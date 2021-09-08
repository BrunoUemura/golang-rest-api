package migrations

import (
	"github.com/BrunoUemura/golang-rest-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}