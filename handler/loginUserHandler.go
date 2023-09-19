package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/dgrijalva/jwt-go"
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

	// JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	secret := os.Getenv("SECRET_KEY")

	token, err := claims.SignedString([]byte(secret))

	if err != nil {
		logger.Errorf("jwt error: %v", passErr.Error())
		sendError(ctx, http.StatusInternalServerError, "erro gerando jwt token.")
		return
	}

	// setting cookie
	ctx.SetCookie("jwt", token, 3600*24, "/", "localhost", false, true)

	// response formatting
	response := loginUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	sendSuccess(ctx, "login-user", response)
}
