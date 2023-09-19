package router

import (
	"github.com/Ticolls/go-auth/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"

	handler.Init()

	v1 := router.Group(basePath)
	{
		v1.GET("/auth", handler.AuthHandler)
	}
}
