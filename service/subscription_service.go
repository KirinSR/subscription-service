package service

import (
	"github.com/google/uuid"
	"kirin.com/subscription-service/models"
	"kirin.com/subscription-service/repositories"
)

type SubscriptionService struct {
	Repo *repositories.SubscriptionRepository
}

func (service *SubscriptionService) CreateSubscription(subscription *models.Subscription) error {
	return service.Repo.Create(subscription)
}

func (service *SubscriptionService) UpdateSubscription(subscription *models.Subscription) error {
	return service.Repo.Update(subscription)
}

func (service *SubscriptionService) DeleteSubscription(subscription *models.Subscription) error {
	return service.Repo.Delete(subscription)
}

func (service *SubscriptionService) GetSubscriptionByUserID(userID uuid.UUID) ([]*models.Subscription, error) {
	return service.Repo.GetByUserID(userID)
}

func (service *SubscriptionService) GetSubscription(id uuid.UUID) (*models.Subscription, error) {
	return service.Repo.GetByID(id)
}

func (service *SubscriptionService) ListSubscriptions() ([]*models.Subscription, error) {
	return service.Repo.List()
}
