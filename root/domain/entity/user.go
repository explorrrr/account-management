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

func ValidateInputPassword(raw_password string) (bool, error) {
	if len(raw_password) < 8 {
		return false, errors.New("Password must be at least 8 characters")
	}

	if len(raw_password) > 128 {
		return false, errors.New("Password must be 128 characters or less")
	}

	return true, nil
}

func ValidateInputUsername(username string) (bool, error) {
	if len(username) > 64 {
		return false, errors.New("Username must be 64 characters or less")
	}

	return true, nil
}

func NewUser(username string, raw_password string) (*UserEntity, error) {

	passwordValidation, passwordErr := ValidateInputPassword(raw_password)
	usernameValidation, usernameErr := ValidateInputUsername(username)

	if passwordValidation == false {
		return &UserEntity{}, passwordErr
	}

	if usernameValidation == false {
		return &UserEntity{}, usernameErr
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(raw_password), 12)

	return &UserEntity{
		Username: username,
		Password: string(password),
	}, nil
}

func (u *UserEntity) AuthUser(inputPassword string) (bool, error) {

	// ハッシュパスワードの比較
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserEntity) IssueJWT(inputPassword string) (string, error) {
	config := config.GetConfig()
	secret := config.GetString("jwt.secret")
	expireSecond := config.GetInt("jwt.expire_second")

	isAuth, err := u.AuthUser(inputPassword)

	if isAuth == false {
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

func (u *UserEntity) ChangePassword(currentPassword string, desiredPassword string) (*UserEntity, error) {

	passwordValidation, passwordErr := ValidateInputPassword(desiredPassword)

	if passwordValidation == false {
		return u, passwordErr
	}

	isAuth, _ := u.AuthUser(currentPassword)

	if isAuth == false {
		return u, errors.New("User is not authenticated")
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(desiredPassword), 12)

	u.Password = string(newPassword)

	return u, nil
}
