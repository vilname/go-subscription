package config

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"subscription-back/docs"
	"subscription-back/internal/controller/rest"
	"subscription-back/internal/util/middleware"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(middleware.EnableCORS)
	router.Use(middleware.Auth)

	router.POST("/subscription/create", rest.CreateSubscription)
	router.GET("/user/:hash", rest.GetUserElement)
	router.POST("/cart/create/:hash", rest.CartCreate)
	router.GET("/card/element/:hash", rest.CardElement)
	router.DELETE("/card/remove/:hash", rest.CardRemove)

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
