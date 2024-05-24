package token

import (
	"errors"
	"github.com/15733012783/weikang/nacos"
	jwt "github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

func CreateToken(str string) (string, error) {
	mySigningKey := []byte(nacos.ApiNac.Jwt.SigningKey)

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		//ID:      "",
		Issuer:  "YuMo",  //发行人
		Subject: "users", //主题
		Audience: jwt.ClaimStrings{
			str,
		}, //受众
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)), //到期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                       //签发日期
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ParseToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(nacos.ApiNac.Jwt.SigningKey), nil
	})

	switch {
	case token.Valid:
		//fmt.Println("You look nice today")
		return true
	case errors.Is(err, jwt.ErrTokenMalformed):
		log.Println("That's not even a token")
		return false
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		log.Println("Invalid signature")
		return false
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		log.Println("Timing is everything")
		return false
	default:
		log.Println("Couldn't handle this token:", err)
		return false
	}
}
