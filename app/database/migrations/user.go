package migrations

import (
	"fmt"

	"github.com/Danangoffic/go-merce/app/models"
	"gorm.io/gorm"
)

// Tips: add UserMigration as a unique Func to register it in migration.go and put db parameter as gorm.DB pointer
// migration will captured depends on Entity structure in models.User

// UserMigration function will run User table migration
func UserMigration(db *gorm.DB) {
	fmt.Println("Run user migration")
	if db.Migrator().HasTable(&models.User{}) {
		db.Migrator().DropTable(&models.User{})
	}
	db.AutoMigrate(&models.User{})
}
