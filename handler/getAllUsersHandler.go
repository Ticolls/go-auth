package handler

import (
	"net/http"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary get all users user
// @Description get all users (need authentication)
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} getAllUsersResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /logout [get]
func GetAllUsersHandler(ctx *gin.Context) {

	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	// response formatting
	response := getAllUsersResponse{}

	for _, u := range users {
		response.Users = append(response.Users, userResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt.Time,
		})
	}

	sendSuccess(ctx, "get-all-users", response)
}
