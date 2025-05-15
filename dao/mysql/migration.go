package mysql

import (
	"cqupt_hub/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&models.Administrator{},
		&models.Department{},
		&models.ImageRecord{},
		&models.Position{},
		&models.Project{},
		&models.Game{},
		&models.Contact{})
}
