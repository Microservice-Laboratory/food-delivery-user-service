package main

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/api"
	"Microservice-Laboratory/food-delivery-user-service/internal/infrastructure/db"
	"log"
	"os"
)

func main() {
	database, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	router := api.SetupRouter(database)

	var host = os.Getenv("HOST")
	var port = os.Getenv("PORT")

	if host == "" {
		host = "0.0.0.0"
	}

	if port == "" {
		port = "8080"
	}

	err = router.Run(host + ":" + port)
	if err != nil {
		panic(err)
	}
}
