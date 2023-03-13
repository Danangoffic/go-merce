package config

import (
	"github.com/Danangoffic/go-merce/app/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// To initialize the GO Application! will return the gorm.DB and gin.Engine pointer
func InitAPP() (*gorm.DB, *gin.Engine) {
	// Set Application routes bootstrap
	routeMapper := routes.RouterStruct{GE: SetupRouter(), DB: InitDB()}
	// load the routers Application in routes folder
	routes.LoadRouters(routeMapper)
	return routeMapper.DB, routeMapper.GE
}
