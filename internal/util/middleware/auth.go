package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscription-back/internal/repository"
)

func Auth(ctx *gin.Context) {
	if ctx.Request.URL.String() == "/subscription/create" {
		ctx.Next()
		return
	}

	hash := ctx.Param("hash")

	subscriptionRepository := repository.NewSubscriptionRepository(ctx)
	id, err := subscriptionRepository.GetByHash(hash)

	if id == 0 || err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
