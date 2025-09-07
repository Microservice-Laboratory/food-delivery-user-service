package service

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/service/users"

	"database/sql"
)

type Services struct {
	UserService *users.UserService
}

func NewServices(db *sql.DB) *Services {
	return &Services{
		UserService: users.NewUserService(db),
	}
}
