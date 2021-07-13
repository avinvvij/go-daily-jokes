package middlewares

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	utils "github.com/avinvvij/go-daily-jokes/utils"
	"fmt"
	"time"
)

func VerifyJWT(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	if(token == ""){
		c.AbortWithStatusJSON(401, gin.H{
			"message": "No auth token",
		})
		return
	}
	claims := utils.NewClaims()
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	fmt.Println(claims.ExpiresAt)
	if(err != nil){
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Invalid auth token",
		})
		return
	}
	if tkn.Valid && ValidateTokenTime(claims.ExpiresAt){
			c.Next()
	}else{
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Invalid auth token",
		})
	}
}

func ValidateTokenTime (timeUnix int64) bool{
	jwtTime := time.Unix(timeUnix, 0)
	currentTime := time.Now()
	return currentTime.Before(jwtTime)
}

func TestMiddleWare(c *gin.Context) {
	authHeader := c.Request.Header.Get("authorization")
	if authHeader != "" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
	}
}
