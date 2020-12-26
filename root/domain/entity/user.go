package entity

import (
	"log"
	"errors"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"account-management/root/config"
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

func (u *UserEntity) AuthUser(inputPassword string) (string, error) {

	config := config.GetConfig()
	secret := config.GetString("jwt.secret")
	expireSecond := config.GetInt("jwt.expire_second")

	// ハッシュパスワードの比較
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))
	if err != nil {
		return "", err
	}

	// jwtトークンの発行
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"iat": time.Now(),
		"exp": time.Now().Add(time.Second * time.Duration(expireSecond)).Unix(),
	})

	token, err := jwtToken.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return token, err
}
