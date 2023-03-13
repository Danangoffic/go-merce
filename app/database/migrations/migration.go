package migrations

import "gorm.io/gorm"

// DoMigration function is to run the migration process that listed inside it
func DoMigration(db *gorm.DB) {
	// table migration here
	UserMigration(db)
	ProductMigration(db)
}
