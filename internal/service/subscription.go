package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"subscription-back/internal/model"
	"subscription-back/internal/repository"
)

type SubscriptionService struct {
	repository *repository.SubscriptionRepository
}

func NewSubscriptionService(ctx *gin.Context) *SubscriptionService {
	return &SubscriptionService{
		repository: repository.NewSubscriptionRepository(ctx),
	}
}

func (service *SubscriptionService) Create(dto model.CreateSubscriptionDto) (model.CreateSubscriptionResult, error) {
	var createSubscriptionResult model.CreateSubscriptionResult

	id, err := service.repository.GetByEmail(dto.Email)

	if err != nil {
		return createSubscriptionResult, err
	}

	if id != 0 {
		return createSubscriptionResult, errors.New("На эту почту уже имеется подписка")
	}

	createSubscriptionResult.Hash, err = service.repository.CreateSubscription(dto.Email)

	if err != nil {
		return createSubscriptionResult, err
	}

	//SendEmail(createSubscriptionResult.Hash)

	return createSubscriptionResult, nil
}
