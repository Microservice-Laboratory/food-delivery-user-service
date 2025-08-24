package main

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api"
	"os"
)

func main() {
	router := api.SetupRouter()

	var host = os.Getenv("HOST")
	var port = os.Getenv("PORT")

	if host == "" {
		host = "0.0.0.0"
	}

	if port == "" {
		port = "8080"
	}

	err := router.Run(host + ":" + port)
	if err != nil {
		panic(err)
	}
}
