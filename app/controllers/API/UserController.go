package API

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Danangoffic/go-merce/app/helpers"
	"github.com/Danangoffic/go-merce/app/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service services.Services
}

// To init UserController Class, pass the *gorm.DB interface
func NewUserController(Service services.Services) *UserController {
	return &UserController{Service: Service}
}

func (u *UserController) GetUserByID(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	userData, err := u.Service.GetUserByID(uint(userID))
	if err != nil {
		response := helpers.ResponseJSON("Failed", http.StatusBadRequest, nil, errors.New("Failed"))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.ResponseJSON("success", http.StatusOK, userData, nil)
	c.JSON(http.StatusOK, response)
	return
}
