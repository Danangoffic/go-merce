package routes

import (
	"github.com/Danangoffic/go-merce/app/controllers"
	Repositories "github.com/Danangoffic/go-merce/app/repositories"
	ServiceApp "github.com/Danangoffic/go-merce/app/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterStruct struct {
	GE *gin.Engine
	DB *gorm.DB
}

func LoadRouters(mapper RouterStruct) {
	repos := populateRepository(mapper.DB)
	// userService :=
	service := ServiceApp.InitService(repos.productRepo)
	pc := controllers.NewProductController(mapper.DB, service)
	userControler := controllers.NewUserController(mapper.DB, service)

	//
	productRoutes := mapper.GE.Group("/products")
	productRoutes.GET("", pc.GetProducts)
	productRoutes.POST("", pc.CreateProduct)

	userRoutes := mapper.GE.Group("/users")
	userRoutes.GET("", userControler.GetUserByID)
}

type RepositoryList struct {
	productRepo Repositories.Repository
}

func populateRepository(db *gorm.DB) RepositoryList {
	productRepo := Repositories.GetRepository(db)
	return RepositoryList{productRepo: productRepo}
}
