package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
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
	_, err := userService.SignUpUser(ctx, username, rawPassword)

	if err != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": err.Error(),
				"code": "9999",
			})
	} else {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": "user create successfully",
				"code": "0000",
			})
	}
}
