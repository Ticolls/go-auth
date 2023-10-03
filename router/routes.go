package router

import (
	"github.com/Ticolls/go-auth/handler"
	"github.com/Ticolls/go-auth/middleware"
	"github.com/gin-gonic/gin"
	docs "github.com/Ticolls/go-auth/docs"
   	swaggerfiles "github.com/swaggo/files"
   	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.Init()
	
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/user", handler.GetAllUsersHandler)
		v1.POST("/register", handler.RegisterUserHandler)
		v1.POST("/login", handler.LoginUserHandler)
		v1.GET("/auth", middleware.AuthMiddleware(), handler.AuthHandler)
		v1.GET("/logout", middleware.AuthMiddleware(), handler.LogoutUserHandler)
	}


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
