package apiv1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func GetDivisionSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting division schedule")

	queryParams := r.URL.Query()
	divisionStr := queryParams.Get("division")

	division, err := strconv.ParseUint(divisionStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid division parameter", http.StatusBadRequest)
		return
	}

	fmt.Println(division)

	response := fmt.Sprintf("Getting schedule for division=%d", division)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func GetRoomSchedule(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	roomStr := queryParams.Get("room")

	room, err := strconv.ParseUint(roomStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid room parameter", http.StatusBadRequest)
		return
	}

	fmt.Println(room)
}

func GetIndividualSchedule(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	individualStr := queryParams.Get("individual")

	individual, err := strconv.ParseUint(individualStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid individual parameter", http.StatusBadRequest)
		return
	}

	fmt.Println(individual)
}

func GetTimestamps(w http.ResponseWriter, r *http.Request) {
	
}


func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/schedule", GetDivisionSchedule).Queries("division", "{division}").Methods("GET")
	r.HandleFunc("/schedule", GetRoomSchedule).Queries("room", "{room}").Methods("GET")
	r.HandleFunc("/schedule", GetIndividualSchedule).Queries("individual", "{individual}").Methods("GET")

	r.HandleFunc("/timestamps", GetTimestamps).Methods("GET")

	http.Handle("/", r)
}