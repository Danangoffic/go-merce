package config

import (
	"github.com/Danangoffic/go-merce/app/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAPP() (*gorm.DB, *gin.Engine) {
	routeMapper := routes.RouterStruct{GE: SetupRouter(), DB: InitDB()}
	routes.LoadRouters(routeMapper)
	return routeMapper.DB, routeMapper.GE
}
