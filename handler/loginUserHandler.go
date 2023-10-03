package handler

import (
	"net/http"

	"github.com/Ticolls/go-auth/schemas"
	"github.com/Ticolls/go-auth/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// @Summary login user
// @Description login and sending jwt in cookies
// @Tags user
// @Accept json
// @Produce json
// @Param request body loginUserRequest true "Request body"
// @Success 200 {object} loginUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
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
	token, err := utils.GenerateToken(user.ID)

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
