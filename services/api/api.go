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

	r.Use(tollbooth_gin.LimitHandler(l))

    SetupRoutes(r)

    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}