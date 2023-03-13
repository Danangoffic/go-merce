package migrations

import (
	"github.com/Danangoffic/go-merce/app/models"
	"gorm.io/gorm"
)

func UserTokenMigration(db *gorm.DB) {
	if db.Migrator().HasTable(&models.UserToken{}) {
		db.Migrator().DropTable(&models.UserToken{})
	}
	db.AutoMigrate(&models.UserToken{})
}
