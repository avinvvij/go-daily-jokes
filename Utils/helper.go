package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	SessionId string `json:"sessionId"`
	jwt.StandardClaims
}

func GenerateJwt(sessionId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"StandardClaims": jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5* time.Minute).Unix(),
		},
		"sessionId": sessionId,
	})	
	signedString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	// return (string)token , nil
	return signedString, err
}

func NewClaims() *Claims{
	claims := &Claims{}
	return claims
}