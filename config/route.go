package config

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"subscription-back/docs"
	"subscription-back/internal/controller/rest"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	//router.Use(middleware.AuthenticationMiddleware)

	router.POST("/subscription/create", rest.CreateSubscription)

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
