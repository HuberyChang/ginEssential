package common

import "github.com/dgrijalva/jwt-go"

var jwtkey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}
