package main

import (
	"encoding/json"
	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody PreviewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserId == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
	}

	// todo: call trip service

	writeJSON(w, http.StatusCreated, "ok")
}
