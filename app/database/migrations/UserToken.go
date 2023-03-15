package migrations

import (
	"fmt"

	"github.com/Danangoffic/go-merce/app/models"
	"gorm.io/gorm"
)

func UserTokenMigration(db *gorm.DB) {
	fmt.Println("Run user token migration")
	if db.Migrator().HasTable(&models.UserToken{}) {
		db.Migrator().DropTable(&models.UserToken{})
	}
	db.AutoMigrate(&models.UserToken{})
}
