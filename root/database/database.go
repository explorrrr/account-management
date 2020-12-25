package database

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"account-management/root/config"
)


type Postgresql struct {
	clientConn *gorm.DB
}

type PostgresqlInterface interface {
	NewClientConnection() *gorm.DB
}

func NewPostgresql() PostgresqlInterface {
	return &Postgresql{}
}

func (postgresql Postgresql) NewClientConnection() *gorm.DB {
	config := config.GetConfig()
	client, err := gorm.Open("postgres", config.GetString("database.dns"))

	if err != nil {
		log.Fatalln("Error in Create client connnection", err)
	}

	return client
}
