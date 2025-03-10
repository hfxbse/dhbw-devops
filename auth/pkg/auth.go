package pkg

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Auth struct {
	SecretKey []byte
}

func (a *Auth) CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(a.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return a.SecretKey, nil
	})
	return err == nil && token.Valid
}
