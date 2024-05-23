package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const APP_KEY = "wanan"

func GenerationToken(userid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userid,
		"exp":  time.Now().Add(time.Hour * time.Duration(10)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	return tokenString, err
}

func AnalysisToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil
	})
	if err != nil {
		return 0, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["user"].(int64), nil
}
