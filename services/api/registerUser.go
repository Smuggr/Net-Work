package api

import (
	"net/http"

	"overseer/data/errors"
	"overseer/data/messages"
	"overseer/data/models"
	"overseer/services/database"

	"github.com/gin-gonic/gin"
)


func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload.Key})
		return
	}

	if err := database.RegisterUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Key})
		return
	}

	tokenString, err := createToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Key})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserRegisterSuccess.Key, "token": tokenString})
}