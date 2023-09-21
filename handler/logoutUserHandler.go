package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutUserHandler(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)

	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "usuário deslogado com sucesso.",
	})
}
