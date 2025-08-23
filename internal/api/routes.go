package api

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api/healthcheck"

	"Microservice-Laboratory/food-delivery-user-service/internal/api/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", healthcheck.HealthCheckHandler)
		v1.GET("/users", user.GetUsersHandler)
	}

	return r
}
