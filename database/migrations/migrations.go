package migrations

import (
	"github.com/diego-arend/boilerplate_go_gin_gorm/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Car{})
	db.AutoMigrate(models.User{})
}
