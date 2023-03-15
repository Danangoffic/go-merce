package migrations

import (
	"fmt"

	"github.com/Danangoffic/go-merce/app/models"
	"gorm.io/gorm"
)

// To run a Product Migration table in database
func ProductMigration(db *gorm.DB) {
	fmt.Println("Run product and product images migration")
	if db.Migrator().HasTable(&models.Product{}) && db.Migrator().HasTable(&models.ProductImages{}) {
		db.Migrator().DropTable(&models.Product{})
		db.Migrator().DropTable(&models.ProductImages{})
	}
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductImages{})
}
