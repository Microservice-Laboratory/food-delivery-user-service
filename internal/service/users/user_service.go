package users

import (
	"Microservice-Laboratory/food-delivery-user-service/internal/domain/entity"
	"Microservice-Laboratory/food-delivery-user-service/internal/service/dto"
	"context"
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (us *UserService) GetUsers(ctx context.Context) ([]dto.UserDto, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: us.db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var users []entity.User
	if err := gormDB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return toUserDtoSlice(users), nil
}

func toUserDtoSlice(users []entity.User) []dto.UserDto {
	userDtos := make([]dto.UserDto, len(users))
	for i, u := range users {
		userDtos[i] = dto.UserDto{
			ID:       u.BaseEntity.ID.String(),
			Username: u.FirstName + " " + u.LastName,
			Email:    u.Email,
		}
	}
	return userDtos
}
