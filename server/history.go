package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type HistoryRequest struct {
	Token   string
	UserID  int
	VideoID int
	Episode int
}

type HistoryResponse struct {
	HistoryList []store.History
}

func RecordHistory(w http.ResponseWriter, r *http.Request) {
	var request HistoryRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("RecordHistory: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	if !store.CheckToken(request.Token) {
		log.Println("RecordHistory: check token failed:")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	err = store.InsertUserHistory(request.UserID, request.VideoID, request.Episode)
	if err != nil {
		log.Println("RecordHistory:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	store.SetCORS(&w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: "Success"})
}

func History(w http.ResponseWriter, r *http.Request) {
	var request HistoryRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("History: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	if !store.CheckToken(request.Token) {
		log.Println("History: check token failed:")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	history, err := store.GetUserHistoryLstMonth(request.UserID)
	if err != nil {
		log.Println("History:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HistoryResponse{HistoryList: history})
}
