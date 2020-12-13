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
	// edinet := new(controller.EdinetController)
	// edinetMaintain := new(controller.MaintainEdinetCompanyController)

	router.GET("/health_check", health.Status)
	// router.GET("/edinet/extract_statement", edinet.GET)
	// router.POST("/edinet/maintain_company", edinetMaintain.POST)

	return router
}
