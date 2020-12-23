package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);not null"`
	Password string `gorm:"type:varchar(128);not null"`
}
