package config

import (
	"github.com/Danangoffic/go-merce/app/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Membuat koneksi ke database
	db, err := setupDatabase()
	if err != nil {
		panic("Failed to connect db")
	}

	// Auto migrate tabel
	database.SeedDB(db)
	return db
}

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
