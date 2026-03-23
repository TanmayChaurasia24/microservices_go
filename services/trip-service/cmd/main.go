package main

import (
	"log"
	"net/http"
	tripHttp "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	mux := http.NewServeMux()

	handler := &tripHttp.HttpHandler{
		Service: svc,
	}

	mux.HandleFunc("POST /preview", handler.HandleTripPreview)

	log.Println("Starting Trip Service API on :8083")
	if err := http.ListenAndServe(":8083", mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
