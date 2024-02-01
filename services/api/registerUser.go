package api

import (
	"net/http"

	"overseer/services/database"
	"overseer/services/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)


func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := database.DB.Where("login = ?", user.Login).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User of this login already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user.Password = string(hashedPassword)

	tokenString, err := createToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	if err := database.RegisterUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "token": tokenString})
}