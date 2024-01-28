package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)


func Initialize(port int) {
    log.Println("Initializing api/v1")
    
	r := gin.Default()
	l := tollbooth.NewLimiter(1, nil)

    r.POST("/authenticate", Authenticate)
	r.POST("/register", Register)

	apiV1Group := r.Group("/api/v1")
	apiV1Group.Use(AuthMiddleware(), tollbooth_gin.LimitHandler(l))

	apiV1Group.GET("/authenticated", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"authenticated": true})
	})

	http.Handle("/", r)

    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}