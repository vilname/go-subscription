package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"subscription-back/internal/repository"
)

func Auth(ctx *gin.Context) {
	if ctx.Request.Method == "OPTIONS" {
		ctx.Writer.WriteHeader(http.StatusOK)
		return
	}

	url := ctx.Request.URL.String()

	if url == "/subscription/create" || strings.Contains(url, "swagger") {
		ctx.Next()
		return
	}

	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	hash := strings.Split(auth, "Bearer ")[1]

	subscriptionRepository := repository.NewSubscriptionRepository(ctx)
	subscriptionModel, err := subscriptionRepository.GetByHash(hash)

	if subscriptionModel.Id == 0 || err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
