package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	id int
	Username string
	Password string
}

func NewUser(username string, raw_password string) (*UserEntity, error) {

	if len(raw_password) < 8 {
		return &UserEntity{}, errors.New("Password must be at least 8 characters")
	}

	if len(raw_password) > 128 {
		return &UserEntity{}, errors.New("Password must be 128 characters or less")
	}

	if len(username) > 64 {
		return &UserEntity{}, errors.New("Username must be 64 characters or less")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(raw_password), 12)

	return &UserEntity{
		Username: username,
		Password: string(password),
	}, nil
}
