package entity

import (
	"github.com/dgrijalva/jwt-go"
	"account-management/root/config"
)

type JWT struct {
	Token string
}

func (t *JWT) Validate() (string, error) {
	config := config.GetConfig()
	secret := config.GetString("jwt.secret")
	_, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
	})

	if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// expire
                return "1001", err
            } else {
				// invalid
                return "9999", err
            }
        }
    }

	// valid
	return "0000", nil
}
