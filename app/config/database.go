package config

import (
	"github.com/Danangoffic/go-merce/app/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB function will initialize the database, run the migration table, seed the database if needed
// and will return as gorm.DB pointer
func InitDB() *gorm.DB {
	// create a database connection
	db, err := setupDatabase()
	if err != nil {
		panic("Failed to connect db")
	}

	// Auto migrate tabel, wait the progress until done with goroutine
	go database.Migrate(db)

	//database seeder
	database.SeedDB(db)
	return db
}

// setupDatabase function will create a connection to database, depends on constant.go, you can change the database configuration.
// This function will return gorm.DB pointer and connection error
func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
