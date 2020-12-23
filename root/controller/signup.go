package controller

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"account-management/root/model"
	_ "github.com/lib/pq"
)

type SignUpController struct {}

func (s SignUpController) POST(c *gin.Context) {
	db, err := gorm.Open("postgres", "postgres://postgres:postgres@account-management-postgres:5432/account_management?sslmode=disable")
	if err != nil {
		log.Fatalln("failed to connect to database", err)
	}
	defer db.Close()
	var user = model.User{Username: "test", Password: "test_password"}
	println(err)

	db.Create(&user)
	db.Save(&user)
	c.String(http.StatusOK, "aaa")
}
