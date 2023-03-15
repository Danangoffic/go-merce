package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

// DoMigration function is to run the migration process that listed inside it
func DoMigration(db *gorm.DB) {
	// table migration here
	fmt.Println("Do migration")
	UserMigration(db)
	UserTokenMigration(db)
	ProductMigration(db)
}
