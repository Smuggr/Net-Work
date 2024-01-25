package api

import (
	// "fmt"
	// "log"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
)




func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/auth")

	apiV1Group := r.Group("/api/v1")
	apiV1Group.Use(AuthMiddleware())

	apiV1Group.GET("/devices", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Resource 1"})
	})

	http.Handle("/", r)
}