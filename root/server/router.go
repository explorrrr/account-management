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
	// edinet := new(controller.EdinetController)

	router.GET("/health_check", health.Status)
	router.POST("/api/sign_up", signUp.POST)
	// router.GET("/edinet/extract_statement", edinet.GET)

	return router
}
