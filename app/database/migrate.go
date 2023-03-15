package database

import (
	"github.com/Danangoffic/go-merce/app/database/migrations"
	"gorm.io/gorm"
)

// Run the database migration
func Migrate(db *gorm.DB) {
	migrations.DoMigration(db)
}
