package service

import (
	"github.com/gin-gonic/gin"
	"subscription-back/internal/model"
	"subscription-back/internal/repository"
	"subscription-back/internal/util/helper"
	"subscription-back/internal/webClient"
)

type CartService struct {
	webClient  *webClient.CartWebCitrus
	repository *repository.SubscriptionRepository
}

func NewCartService(ctx *gin.Context) *CartService {
	return &CartService{
		webClient:  webClient.NewCartWebCitrus(),
		repository: repository.NewSubscriptionRepository(ctx),
	}
}

func (cartService *CartService) Create(dto model.CreateCartDto, hash string) error {
	subscription, err := cartService.repository.GetByHash(hash)
	if err != nil {
		return err
	}

	dto.Email = subscription.Email
	dto.PhoneNumber = helper.ClearPhone(dto.PhoneNumber)

	webCitrusSuccessResponse, err := cartService.webClient.Create(dto)
	if err != nil {
		return err
	}

	err = cartService.repository.UpdateCardUuid(hash, webCitrusSuccessResponse.Data.SerialNumber)
	if err != nil {
		return err
	}

	return nil
}

func (cartService *CartService) GetElement(dto model.CardElementDto) (model.CardElementResult, error) {
	var cardResult model.CardElementResult

	subscription, err := cartService.repository.IsExistCard(dto.Hash)

	if err != nil {
		return cardResult, err
	}

	_ = subscription

	return cardResult, nil
}
