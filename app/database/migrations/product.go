package migrations

import (
	"github.com/Danangoffic/go-merce/app/models"
	"gorm.io/gorm"
)

// To run a Product Migration table in database
func ProductMigration(db *gorm.DB) {
	if db.Migrator().HasTable(&models.Product{}) {
		db.Migrator().DropTable(&models.Product{})
	}
	db.AutoMigrate(&models.Product{})
}
