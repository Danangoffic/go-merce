package config

import (
	"time"

	"github.com/Danangoffic/go-merce/app/database"
	"github.com/Danangoffic/go-merce/app/routes"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitAPP() (*gorm.DB, *gin.Engine) {
	routeMapper := routes.RouterStruct{GE: setupRouter(), DB: initDB()}
	routes.LoadRouters(routeMapper)
	return routeMapper.DB, routeMapper.GE
}

func initDB() *gorm.DB {
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

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(limits.RequestSizeLimiter(60))
	return router
}
