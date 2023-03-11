package routes

import (
	"github.com/Danangoffic/go-merce/app/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRouters(gi *gin.Engine, db *gorm.DB) {
	pc := controllers.NewProductController(db)

	productRoutes := gi.Group("/products")
	productRoutes.GET("", pc.GetProducts)
}
