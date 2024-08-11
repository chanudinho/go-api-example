package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if err != nil || !token.Valid || !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		fmt.Println(claims["role"], role)
		if role != "" && claims["role"] != role {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		userID := claims["userID"].(float64)

		c.Set("userID", uint(userID))
		c.Set("role", claims["role"])

		c.Next()
	}
}
