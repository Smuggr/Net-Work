package apiv1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	scraper "overseer/services/scraper/common"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func Initialize(port int) {
    fmt.Println("Initializing api/v1")
    
    scraper.Initialize()

    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}