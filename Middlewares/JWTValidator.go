package middlewares

import (
	utils "github.com/avinvvij/go-daily-jokes/Utils"
	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func VerifyJWT(c *gin.Context, token string) {
	claims := utils.NewClaims()
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	if(err != nil){
		
	}
}
