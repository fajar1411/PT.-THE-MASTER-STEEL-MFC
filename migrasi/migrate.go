package migrasi

import (
	"master/domain/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Employee{})
	db.AutoMigrate(&model.Salary{})

}
