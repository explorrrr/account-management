package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	id int
	username string
	password string
}

func NewUser(id int, username string, raw_password string) (*UserEntity, error) {

	// usernameとpasswordの入力チェック(文字数やパスワード強度など)

	// passwordのハッシュ化を入れる
	password _ := bcrypt.GenerateFromPassword([]byte(raw_password), 12)

	return &UserEntity{
		username: username,
		password: password,
	}, nil
}
