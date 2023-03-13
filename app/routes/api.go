package routes

import (
	"github.com/Danangoffic/go-merce/app/controllers/API"
	ServiceApp "github.com/Danangoffic/go-merce/app/services"
)

// Load API function will populate the API resources
func LoadAPI(mapper RouterStruct, service ServiceApp.Services) {
	// userControler := API.NewUserController(service)

	apiRoute := mapper.GE.Group("/api")
	//
	productRoutes := apiRoute.Group("/products")

	pc := API.NewProductController(service)
	// load api resource by saved variable
	productRoutes.GET("", pc.GetProducts)

	// load api resource by lazy load controller
	productRoutes.POST("", API.NewProductController(service).CreateProduct)

	userRoutes := apiRoute.Group("/users")
	userRoutes.GET("", API.NewUserController(service).GetUserByID)
}
