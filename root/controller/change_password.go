package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"account-management/root/domain/service"
)

type ChangePasswordRequest struct {
	Username string `json:"username"`
	CurrentPassword string `json:"current_password"`
	DesiredPassword string `json:"desired_password"`
}

type ChangePasswordController struct {}

func (a ChangePasswordController) POST(ctx *gin.Context) {

	var request ChangePasswordRequest
	ctx.BindJSON(&request)
	username := request.Username
	currentPassword := request.CurrentPassword
	desiredPassword := request.DesiredPassword

	userService := service.UserService{}
	result, code, err := userService.ChangePassword(ctx, username, currentPassword, desiredPassword)

	if result == false {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code": code,
				"msg": err.Error(),
			})
	} else {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code": code,
				"msg": "Password changed successfully",
			})
	}
}
