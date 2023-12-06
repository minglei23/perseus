package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type PointsRequest struct {
	Token  string
	UserID int
}

type PointsResponse struct {
	Checked bool
	Points  int
}

func Points(w http.ResponseWriter, r *http.Request) {
	var request PointsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Points: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	if !store.CheckToken(request.Token) {
		log.Println("Points: check token failed:")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	points, err := store.GetPoints(request.UserID)
	if err != nil {
		log.Println("Points:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	store.SetCORS(&w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PointsResponse{Points: points})
}

func AlreadyCheckin(w http.ResponseWriter, r *http.Request) {
	var request PointsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("AlreadyCheckin: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	if !store.CheckToken(request.Token) {
		log.Println("AlreadyCheckin: check token failed:")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	checked, err := store.GetIfUserCheckedToday(request.UserID)
	if err != nil {
		log.Println("AlreadyCheckin:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	store.SetCORS(&w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PointsResponse{Checked: checked})
}

func Checkin(w http.ResponseWriter, r *http.Request) {
	var request PointsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Checkin: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	if !store.CheckToken(request.Token) {
		log.Println("Checkin: check token failed:")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	checked, err := store.GetIfUserCheckedToday(request.UserID)
	if err != nil {
		log.Println("Checkin: GetIfUserCheckedToday:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	points := 0
	if !checked {
		if err := store.InsertActivities(request.UserID, 1, 1); err != nil {
			log.Println("Checkin: InsertActivities:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		points, err = store.GetPoints(request.UserID)
		if err != nil {
			log.Println("Checkin: GetPoints:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		points += 1
		if err := store.UpdateUserPoints(request.UserID, points); err != nil {
			log.Println("Checkin: UpdateUserPoints:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	store.SetCORS(&w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PointsResponse{Checked: true, Points: points})
}
