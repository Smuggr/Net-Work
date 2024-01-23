package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func GetDivisionSchedule(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting division schedule")

	queryParams := r.URL.Query()
	divisionStr := queryParams.Get("division")

	division, err := strconv.ParseUint(divisionStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid division parameter", http.StatusBadRequest)
		return
	}

	log.Println(division)

	response := fmt.Sprintf("Getting schedule for division=%d", division)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}


func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/schedule", GetDivisionSchedule).Queries("division", "{division}").Methods("GET")

	http.Handle("/", r)
}