package main

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api"
	"os"
)

func main() {
	router := api.SetupRouter()

	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run("localhost:" + port)
}
