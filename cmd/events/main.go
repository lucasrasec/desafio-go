package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasrasec/desafio-go/internal/handlers"
	"github.com/lucasrasec/desafio-go/internal/usecase"
)

func main() {
	usecase := usecase.NewEventUseCases()
	if err := usecase.LoadData("../../data.json"); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	handler := handlers.NewEventHandler(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/events", handler.GetEvents).Methods("GET")
	r.HandleFunc("/events/{eventId}", handler.GetEventByID).Methods("GET")
	r.HandleFunc("/events/{eventId}/spots", handler.GetEventSpots).Methods("GET")
	r.HandleFunc("/events/{eventId}/reserve", handler.ReserveSpot).Methods("POST")

	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
