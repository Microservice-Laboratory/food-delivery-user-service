package user

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/service/users"

	"github.com/gin-gonic/gin"
)

type UserEndpoint struct {
	userService *users.UserService
}

func NewUserEndpoint(userService *users.UserService) *UserEndpoint {
	return &UserEndpoint{userService: userService}
}

func (ue *UserEndpoint) GetUsersHandler(c *gin.Context) {
	users, err := ue.userService.GetUsers(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}
