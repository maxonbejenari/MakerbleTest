package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/maxonbejenari/testWebApp/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	// function will run before the request reaches our handler
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization") // retrieves the Authorization header from the incoming HTTP Req
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization header",
			})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer") // removes bearer prefix from the header
		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}
		// if the token is valid, the middleware adds the user's info
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next() // this tells Gin to continue processing the request
	}
}
