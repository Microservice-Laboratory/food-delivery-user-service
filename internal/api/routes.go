package api

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api/healthcheck"

	"Microservice-Laboratory/food-delivery-user-service/internal/api/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthcheck.HealthCheckHandler)
	r.GET("/users", user.GetUsersHandler)

	return r
}
