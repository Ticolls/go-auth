package middleware

import (
	"net/http"

	"github.com/Ticolls/go-auth/config"
	"github.com/Ticolls/go-auth/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	logger := config.GetLogger("middleware")

	return func(ctx *gin.Context) {

		cookie, err := ctx.Cookie("jwt")

		if err != nil {
			logger.Errorf("cookie error: %v", err)
			ctx.Header("content-type", "application/json")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message":   "você não tem permissão para entrar aqui.",
				"errorCode": http.StatusUnauthorized,
			})
			return
		}

		id, err := utils.ValidateToken(cookie)

		if err != nil {
			logger.Errorf("jwt error: %v", err)
			ctx.Header("content-type", "application/json")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message":   "você não tem permissão para entrar aqui.",
				"errorCode": http.StatusUnauthorized,
			})
			return
		}

		ctx.Set("id", id)
		ctx.Next()
	}
}
