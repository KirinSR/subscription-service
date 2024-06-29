package service

import (
	"github.com/google/uuid"
	"kirin.com/subscriptio-service/models"
	"kirin.com/subscriptio-service/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (service *UserService) CreateNewUser(user *models.User) error {
	return service.Repo.Create(user)
}

func (service *UserService) UpdateUser(user *models.User) error {
	return service.Repo.Update(user)
}

func (service *UserService) DeleteUser(id uuid.UUID) (*models.User, error) {
	return service.Repo.deleteUser(id)
}

func (service *UserService) GetByID(id uuid.UUID) (*models.User, error) {
	return service.Repo.GetByID(id)
}

func (service *UserService) GetAllUsers([]*models.User) error {
	return service.Repo.List()
}
