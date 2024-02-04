package api

import (
	"net/http"

	"overseer/services/database"
	"overseer/services/models"

	"github.com/gin-gonic/gin"
)


func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	
	if err := database.UpdateUser(database.DB, user.Login, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User updated successfully"})
}