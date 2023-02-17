package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type author interface {
	GenToken()
	ParseToken()
}

type jwtToken struct {
}

type DefaultClaims struct {
	jwt.RegisteredClaims
	UserId string
}

func CreateToken(signKey string) (string, error) {
	c := DefaultClaims{UserId: "123"}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	return token.SignedString(signKey)
}

func ParserToken(tokenStr, signKey string) (*DefaultClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenStr, &DefaultClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	return token.Claims.(*DefaultClaims), err
}
