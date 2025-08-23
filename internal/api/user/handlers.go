package user

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []dto.UserDto{
	{ID: "123", Username: "Username", Email: "Email"},
}

func GetUsersHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
