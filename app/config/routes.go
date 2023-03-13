package config

import (
	"time"

	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

// SetupRouter function will set the CORS a.k.a "Cross-origin resource sharing".
// Limit request API by default is 60, you can change it, be careful.
// Router use cors and limits middleware from gin-contrib.
// This function will return a gin.Engine pointer
func SetupRouter() *gin.Engine {
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
