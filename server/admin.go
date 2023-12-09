package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type UploadVideoInfoRequest struct {
	Admin       string
	Name        string
	Type        int
	TotalNumber int
	BaseUrl     string
}

type UploadVideoInfoResponse struct {
	ID int
}

func UploadVideoInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request UploadVideoInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("UploadVideoInfo: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if request.Admin != store.Admin {
		log.Println("UploadVideoInfo: invalid admin credentials.")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id, err := store.InsertVideo(request.Name, request.Type, request.TotalNumber, request.BaseUrl)
	if err != nil {
		log.Println("UploadVideoInfo: error inserting video:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UploadVideoInfoResponse{ID: id})
}
