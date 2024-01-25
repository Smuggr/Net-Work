package api

import (
	// "fmt"
	// "log"
	"net/http"

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

func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}