package api

import (
	"net/http"

	"overseer/data/errors"
	"overseer/data/messages"
	"overseer/data/models"
	"overseer/services/database"

	"github.com/gin-gonic/gin"
)


func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload.Error()})
		return
	}
	
	if err := database.UpdateUser(database.DB, user.Login, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserUpdateSuccess.Msg()})
}