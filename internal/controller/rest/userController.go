package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscription-back/internal/model"
)

// GetUserElement godoc
// @Tags Пользователи
// @Summary Детальная страница
// @Param hash path string true "Hash"
// @Accept  json
// @Produce  json
//
// @Success 200	{object} model.UserElementResult
// @Failure	500	{object} helper.ErrorResponse "Другие ошибки"
// @Router /user/{hash} [get]
func GetUserElement(ctx *gin.Context) {
	hash := ctx.Param("hash")
	_ = hash
	
	var userElementResult model.UserElementResult

	ctx.JSON(http.StatusOK, userElementResult)
}
