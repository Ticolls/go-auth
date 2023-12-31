package router

import (
	"os"

	"github.com/Ticolls/go-auth/config"
	"github.com/Ticolls/go-auth/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {

	logger := config.GetLogger("router")

	router := gin.Default()

	// setting cors
	router.Use(middleware.CORSMiddleware())

	initializeRoutes(router)

	// Setting the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	err := router.Run(":" + port)

	if err != nil {
		logger.Errorf("Error running the server: %v", err)
		return
	}
}
