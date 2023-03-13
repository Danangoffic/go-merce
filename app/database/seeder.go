package database

import (
	"github.com/Danangoffic/go-merce/app/database/seeders"
	"gorm.io/gorm"
)

// The SeedDB function will run the seeders on the tables in the database.
// Just add seeders file in seeders folders to keep it clean
// then call the function and pass the db parameter
func SeedDB(db *gorm.DB) {
	seeders.RunSeederUser(db)
	seeders.RunSeederProduct(db)
}
