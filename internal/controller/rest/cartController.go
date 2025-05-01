package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"subscription-back/internal/model"
	"subscription-back/internal/service"
	"subscription-back/internal/util/helper"
)

// CartCreate godoc
// @Tags Карта
// @Param hash path string true "Hash"
// @Param request body model.CreateCartDto true "Создание карты"
// @Summary Создание карты
// @Accept json
// @Produce json
//
// @Failure	400	{object} helper.ErrorValidate "Ошибка валидации"
// @Failure	500	{object} helper.ErrorResponse "Другие ошибки"
// @Router /cart/create/{hash} [post]
func CartCreate(ctx *gin.Context) {
	var createCartDto model.CreateCartDto

	hash := ctx.Param("hash")

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	if err := json.Unmarshal(body, &createCartDto); err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	cartService := service.NewCartService(ctx)
	err = cartService.Create(createCartDto, hash)
	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ctx)
}

// CardElement godoc
// @Tags Карта
// @Summary Детальная страница
// @Param hash path string true "Hash"
// @Accept  json
// @Produce  json
//
// @Success 200	{object} model.CardElementResult
// @Failure	500	{object} helper.ErrorResponse "Другие ошибки"
// @Router /card/element/{hash} [get]
func CardElement(ctx *gin.Context) {
	var cardDto model.CardElementDto

	cardDto.Hash = ctx.Param("hash")

	cartService := service.NewCartService(ctx)
	cardElement, err := cartService.GetElement(cardDto)
	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, cardElement)
}

// CardRemove godoc
// @Tags Карта
// @Param hash path string true "Hash"
// @Summary Удаление карты
// @Accept json
// @Produce json
// @Failure	400	{object} helper.ErrorValidate "Ошибка валидации"
// @Failure	500	{object} helper.ErrorResponse "Другие ошибки"
// @Router /cart/remove/{hash} [delete]
func CardRemove(ctx *gin.Context) {
	cartService := service.NewCartService(ctx)
	err := cartService.CardRemove(ctx.Param("hash"))
	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ctx)
}
