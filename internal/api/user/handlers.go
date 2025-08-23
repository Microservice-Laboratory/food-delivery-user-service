package user

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []dto.UserDto{
	{ID: "123", Username: "john_doe", Email: "john.doe@example.com"},
}

func GetUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
