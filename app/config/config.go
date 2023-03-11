package config

import (
	"github.com/Danangoffic/go-merce/app/database"
	"github.com/gin-gonic/gin"
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
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/ecommerce_go?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
}
