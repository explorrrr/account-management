package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"account-management/root/domain/service"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthController struct {}

func (a AuthController) POST(ctx *gin.Context) {

	var request AuthRequest
	ctx.BindJSON(&request)
	username := request.Username
	rawPassword := request.Password

	userService := service.UserService{}
	jwtToken, err := userService.AuthUser(ctx, username, rawPassword)

	if err != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"token": "",
				"code": "9999",
			})
	} else {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"token": jwtToken.Token,
				"code": "0000",
			})
	}
}
