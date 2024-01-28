package api

import (
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"

	"overseer/services/database"
	"overseer/services/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func createToken(username string) (string, error) {
	jwt_token_lifespan, err := strconv.Atoi(os.Getenv("API_JWT_TOKEN_LIFESPAN_MINUTES"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Duration(jwt_token_lifespan) * time.Minute).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := strings.Split(header, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_TOKEN")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Authenticate(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := database.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := createToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}