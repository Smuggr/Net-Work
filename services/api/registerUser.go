package api

import (
	"net/http"

	"overseer/services/database"
	"overseer/services/errors"
	"overseer/services/models"

	"github.com/gin-gonic/gin"
)


func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload.Error()})
		return
	}

	if err := database.RegisterUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := createToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "token": tokenString})
}