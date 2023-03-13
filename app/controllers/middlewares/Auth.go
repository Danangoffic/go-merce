package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Danangoffic/go-merce/app/helpers"
	"github.com/Danangoffic/go-merce/app/services"
	"github.com/gin-gonic/gin"
)

func ValidateAuthBearer(service services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.ResponseJSON("Authentication Failed", http.StatusUnauthorized, nil, errors.New("Authentication Failed!"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayAuthHeader := strings.Split(authHeader, " ")
		if len(arrayAuthHeader) == 2 {
			tokenString = arrayAuthHeader[1]
		}

		token := tokenString
		user, err := service.GetUserTokenData(token)
		if err != nil {
			response := helpers.ResponseJSON("Authentication Failed", http.StatusUnauthorized, nil, errors.New("Authentication Failed!"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("user", user)
	}
}
