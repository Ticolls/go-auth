package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary logout user
// @Description logout cleaning the jwt cookie
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} loginUserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /logout [get]
func LogoutUserHandler(ctx *gin.Context) {
	_, err := ctx.Cookie("jwt")

	if err != nil {
		sendError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)

	sendSuccess(ctx, "logout-user", nil)
}
