package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func Initialize(port int) {
    log.Println("Initializing api/v1")
    
	r := gin.Default()
    SetupRoutes(r)

    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}