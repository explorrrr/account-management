package server

import (
	"github.com/gin-gonic/gin"
	"account-management/root/controller"
)

// NewRouter returns router
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controller.HealthController)
	signUp := new(controller.SignUpController)
	auth := new(controller.AuthController)
	validateToken := new(controller.ValidateTokenController)

	router.GET("/health_check", health.Status)
	router.POST("/api/sign_up", signUp.POST)
	router.POST("/api/auth", auth.POST)
	router.POST("/api/validate_token", validateToken.POST)

	return router
}
