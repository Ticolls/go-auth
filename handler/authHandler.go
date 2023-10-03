package handler

import (
	"net/http"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary get all logged user info
// @Description need authentication, first create a account and log in
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} userResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth [get]
func AuthHandler(ctx *gin.Context) {

	id, exists := ctx.Get("id")

	if !exists {
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
