package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "auth",
	})
}
