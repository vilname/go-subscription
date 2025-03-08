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
	var cartElementResult model.CardElementResult

	serialNumber, err := cartService.repository.GetSerialNumber(dto.Hash)

	if err != nil {
		return cartElementResult, err
	}

	if serialNumber == "" {
		return cartElementResult, err
	}

	webCitrusSuccessResponse, err := cartService.webClient.GetElement(serialNumber)

	if err != nil {
		return cartElementResult, err
	}

	cartElementResult.FirstName = webCitrusSuccessResponse.Data.Fields.FirstName
	cartElementResult.LastName = webCitrusSuccessResponse.Data.Fields.LastName
	cartElementResult.PhoneNumber = webCitrusSuccessResponse.Data.Fields.PhoneNumber
	cartElementResult.Email = webCitrusSuccessResponse.Data.Fields.Email
	cartElementResult.BirdDate = webCitrusSuccessResponse.Data.Fields.BirdDate
	cartElementResult.Link = webCitrusSuccessResponse.Data.Link
	cartElementResult.PassNumber = webCitrusSuccessResponse.Data.PassNumber

	return cartElementResult, nil
}

func (cartService *CartService) CardRemove(hash string) error {

	serialNumber, err := cartService.repository.GetSerialNumber(hash)
	if err != nil {
		return err
	}

	err = cartService.webClient.RemoveCardClient(serialNumber)
	if err != nil {
		return err
	}

	return nil
}
