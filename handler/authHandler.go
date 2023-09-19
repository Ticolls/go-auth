package handler

import (
	"net/http"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/Ticolls/go-auth/utils"
	"github.com/gin-gonic/gin"
)

func AuthHandler(ctx *gin.Context) {
	cookie, err := ctx.Cookie("jwt")

	if err != nil {
		logger.Errorf("cookie error: %v", err)
		sendError(ctx, http.StatusUnauthorized, "você não tem permissão para entrar aqui.")
		return
	}

	id, err := utils.ValidateToken(cookie)

	if err != nil {
		logger.Errorf("jwt error: %v", err)
		sendError(ctx, http.StatusUnauthorized, "você não tem permissão para entrar aqui.")
		return
	}

	user := schemas.User{}

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		logger.Errorf("user not found: %v", err)
		sendError(ctx, http.StatusInternalServerError, "usuário não encontrado.")
		return
	}

	response := userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
	}

	sendSuccess(ctx, "auth-handler", response)

}
