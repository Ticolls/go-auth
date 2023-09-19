package handler

import (
	"net/http"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserHandler(ctx *gin.Context) {
	request := loginUserRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	var emailErr, passErr error

	if emailErr = db.Where("email = ?", request.Email).First(&user).Error; emailErr != nil {
		logger.Errorf("usuário não encontrado: %v", passErr.Error())
	}

	if passErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); passErr != nil {
		logger.Errorf("Senha incorreta: %v", passErr.Error())
	}

	if emailErr != nil || passErr != nil {
		sendError(ctx, http.StatusBadRequest, "Email ou senha inválidos.")
		return
	}

	sendSuccess(ctx, "login-user", user)
}
