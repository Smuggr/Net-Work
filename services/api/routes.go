package api

import (
	// "fmt"
	// "log"
	"net/http"

	"overseer/services/database"
	"overseer/services/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine) {
	r.POST("/authenticate", Authenticate)


	apiV1Group := r.Group("/api/v1")
	apiV1Group.Use(AuthMiddleware())

	apiV1Group.GET("/authenticated", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"authenticated": true})
	})

	http.Handle("/", r)
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

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	return token.SignedString([]byte("amongusballs"))
}

func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}