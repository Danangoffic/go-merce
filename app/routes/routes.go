package routes

import (
	"github.com/Danangoffic/go-merce/app/controllers"
	repoProduct "github.com/Danangoffic/go-merce/app/repositories/product"
	"github.com/Danangoffic/go-merce/app/services/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterStruct struct {
	GE *gin.Engine
	DB *gorm.DB
}

func LoadRouters(mapper RouterStruct) {
	repos := populateRepository(mapper.DB)
	productService := product.InitService(repos.productRepo)
	pc := controllers.NewProductController(mapper.DB, productService)

	productRoutes := mapper.GE.Group("/products")
	productRoutes.GET("", pc.GetProducts)
}

type RepositoryList struct {
	productRepo repoProduct.Repository
}

func populateRepository(db *gorm.DB) RepositoryList {
	productRepo := repoProduct.GetRepository(db)
	return RepositoryList{productRepo: productRepo}
}
