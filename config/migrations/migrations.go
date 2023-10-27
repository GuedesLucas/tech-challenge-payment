package migrations

import (
	"tech-challenge-payment/internal/payment/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Payment{})
}

func RollbackMigrations(db *gorm.DB) {
	db.Migrator().DropTable(&models.Payment{})
}
