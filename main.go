package main

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api"
)

func main() {
	router := api.SetupRouter()
	router.Run("localhost:8080")
}
