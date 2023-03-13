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

// LoadRouters function will load all Application routes and pass the mapper as RouterStruct in Application bootstrap
func LoadRouters(mapper RouterStruct) {
	// populate the repositories
	populateRepos := populateRepository(mapper.DB)
	// userService :=
	service := ServiceApp.InitService(populateRepos.repos)
	pc := controllers.NewProductController(service)
	userControler := controllers.NewUserController(service)

	//
	productRoutes := mapper.GE.Group("/products")
	productRoutes.GET("", pc.GetProducts)
	productRoutes.POST("", pc.CreateProduct)

	userRoutes := mapper.GE.Group("/users")
	userRoutes.GET("", userControler.GetUserByID)
}

type RepositoryList struct {
	repos Repositories.Repository
}

// populateRepository function will populate all the repositories that belong to the application
func populateRepository(db *gorm.DB) RepositoryList {
	repos := Repositories.GetRepository(db)
	return RepositoryList{repos: repos}
}
