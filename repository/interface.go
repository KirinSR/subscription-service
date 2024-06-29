package repository

import "kirin.com/subscription-service/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
}
