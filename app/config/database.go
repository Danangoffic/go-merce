package config

import (
	"time"

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
		panic(err.Error())
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1000)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

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
