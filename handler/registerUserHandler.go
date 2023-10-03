package handler

import (
	"net/http"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// @Summary Register user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param request body registerUserRequest true "Request body"
// @Success 200 {object} registerUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func RegisterUserHandler(ctx *gin.Context) {

	request := registerUserRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//Hashpassword logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	if err != nil {
		logger.Errorf("error hashing the password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// response formatting
	response := registerUserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}

	sendSuccess(ctx, "register-user", &response)

}
