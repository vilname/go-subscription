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

// CreateSubscription godoc
// @Tags Подписка
// @Param request body model.CreateSubscriptionDto true "Подписка"
// @Summary Создание подписки
// @Accept json
// @Produce json
// @Failure	400	{object} helper.ErrorValidate "Ошибка валидации"
// @Failure	500	{object} helper.ErrorResponse "Другие ошибки"
// @Router /subscription/create [post]
func CreateSubscription(ctx *gin.Context) {
	var subscriptionDto model.CreateSubscriptionDto

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	if err := json.Unmarshal(body, &subscriptionDto); err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	subscriptionService := service.NewSubscriptionService(ctx)
	createSubscriptionResult, err := subscriptionService.Create(subscriptionDto)

	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, createSubscriptionResult)
}
