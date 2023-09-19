package handler

import (
	"github.com/Ticolls/go-auth/schemas"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {

	request := registerUserRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		return
	}

	//Hashpassword logic

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		return
	}

	sendSuccess(ctx, "register-user", user)

}
