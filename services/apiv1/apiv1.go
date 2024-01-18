package apiv1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
    scraper "overseer/services/scraper/common"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func Initialize() {
    fmt.Println("Initializing api/v1")
    
    scraper.Initialize()

    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}