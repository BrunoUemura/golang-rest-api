package migrations

import (
	"github.com/BrunoUemura/golang-rest-api/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}