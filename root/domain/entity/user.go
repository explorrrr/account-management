package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	id int
	Username string
	Password string
}

func NewUser(username string, raw_password string) (*UserEntity, error) {

	// usernameとpasswordの入力チェック(文字数やパスワード強度など)

	password, _ := bcrypt.GenerateFromPassword([]byte(raw_password), 12)

	return &UserEntity{
		Username: username,
		Password: string(password),
	}, nil
}
