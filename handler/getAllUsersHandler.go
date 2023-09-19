package handler

import (
	"net/http"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/gin-gonic/gin"
)

func GetAllUsersHandler(ctx *gin.Context) {

	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	// response formatting
	response := []userResponse{}

	for _, u := range users {
		response = append(response, userResponse{
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
