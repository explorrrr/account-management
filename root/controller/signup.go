package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"account-management/root/domain/service"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpController struct {}

func (s SignUpController) POST(ctx *gin.Context) {

	var request SignUpRequest
	ctx.BindJSON(&request)
	username := request.Username
	rawPassword := request.Password

	userService := service.UserService{}
	result := userService.SignUpUser(ctx, username, rawPassword)

	if result == false {
		ctx.String(http.StatusOK, "user already exist, please try another username.")
	} else {
		ctx.String(http.StatusOK, "user create successfully")
	}
}
