package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscription-back/internal/model"
)

// GetCabinetElement godoc
// @Tags Личный кабинет
// @Summary Личный кабинет пользователя
// @Param hash path string true "Hash"
// @Accept  json
// @Produce  json
//
// @Success 200	{object} model.UserElementResult
// @Failure	500	{object} helper.ErrorResponse "Другие ошибки"
// @Router /cabinet/{hash} [get]
func GetCabinetElement(ctx *gin.Context) {
	hash := ctx.Param("hash")
	_ = hash

	var userElementResult model.UserElementResult

	ctx.JSON(http.StatusOK, userElementResult)
}
