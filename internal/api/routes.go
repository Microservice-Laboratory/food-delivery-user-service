package api

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api/healthcheck"
	"Microservice-Laboratory/food-delivery-user-service/internal/api/user"
	"Microservice-Laboratory/food-delivery-user-service/internal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(database *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(SetJSONContentType())

	services := service.NewServices(database)
	userEndpoint := user.NewUserEndpoint(services.UserService)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", healthcheck.HealthCheckHandler)
		v1.GET("/users", userEndpoint.GetUsersHandler)
	}

	return r
}
