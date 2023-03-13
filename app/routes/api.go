package routes

import (
	"github.com/Danangoffic/go-merce/app/controllers/API"
	ServiceApp "github.com/Danangoffic/go-merce/app/services"
)

// Load API function will populate the API resources
func LoadAPI(mapper RouterStruct, service ServiceApp.Services) {
	pc := API.NewProductController(service)
	userControler := API.NewUserController(service)

	apiRoute := mapper.GE.Group("/api")
	//
	productRoutes := apiRoute.Group("/products")
	productRoutes.GET("", pc.GetProducts)
	productRoutes.POST("", pc.CreateProduct)

	userRoutes := apiRoute.Group("/users")
	userRoutes.GET("", userControler.GetUserByID)
}
