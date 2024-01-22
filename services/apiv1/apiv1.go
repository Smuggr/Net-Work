package apiv1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"overseer/services/scraper/common"

	"github.com/gorilla/mux"
)


func Initialize(port int) {
    fmt.Println("Initializing api/v1")
    
    scraper.Initialize()

    r := mux.NewRouter()
    SetupRoutes(r)

    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}