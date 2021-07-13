package middlewares

import (
	"github.com/gin-gonic/gin"
)

func VerifyJWT(c *gin.Context, token string) {
	// claims := utils.NewClaims()
	// tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return "", nil
	// })
	// if(err != nil){
	//
	// }
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
