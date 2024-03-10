package api

import (
	"net/http"
	"os"
	"strings"

	"network/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized})
			c.Abort()
			return
		}

		// "Bearer <token>"
		tokenString := strings.Split(header, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_TOKEN")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidToken})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		login, ok := claims["login"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidTokenFormat})
			c.Abort()
			return
		}

		c.Set("login", login)
		c.Next()
	}
}

func DeviceAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized})
			c.Abort()
			return
		}

		// "Bearer <token>"
		tokenString := strings.Split(header, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_TOKEN")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidToken})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		client_id, ok := claims["client_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidTokenFormat})
			c.Abort()
			return
		}

		c.Set("client_id", client_id)
		c.Next()
	}
}
