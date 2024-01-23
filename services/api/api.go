package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func Initialize(port int) {
    log.Println("Initializing api/v1")
    
    r := mux.NewRouter()
    SetupRoutes(r)

    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}