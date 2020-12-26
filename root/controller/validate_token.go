package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"account-management/root/domain/service"
)

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenController struct {}

func (a ValidateTokenController) POST(ctx *gin.Context) {

	var request ValidateTokenRequest
	ctx.BindJSON(&request)
	token := request.Token

	userService := service.UserService{}
	code, err := userService.ValidateToken(ctx, token)

	if err != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code": code,
			})
	} else {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code": code,
			})
	}
}
