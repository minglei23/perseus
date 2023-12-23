package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type EpisodesRequest struct {
	Token  string
	UserID int
}

type UnlockEpisodeRequest struct {
	Token   string
	UserID  int
	VideoID int
	Episode int
}

type EpisodesResponse struct {
	VideoEpisodeList []store.UserVideoEpisode
}

func UnlockEpisode(w http.ResponseWriter, r *http.Request) {
	var request UnlockEpisodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, "Invalid request format", http.StatusBadRequest, err)
		return
	}
	if !store.CheckToken(request.Token) {
		respondWithError(w, "Invalid request format", http.StatusBadRequest, nil)
		return
	}
	exist, err := store.UserEpisodeExist(request.UserID, request.VideoID, request.Episode)
	if err != nil {
		respondWithError(w, "Internal server error", http.StatusBadRequest, err)
		return
	}
	if exist {
		respondWithError(w, "Episode has been unlocked", http.StatusBadRequest, nil)
		return
	}
	points, err := store.GetPoints(request.UserID)
	if err != nil {
		respondWithError(w, "Internal server error", http.StatusBadRequest, err)
		return
	}
	if points < 1 {
		respondWithError(w, "User points are not enough", http.StatusBadRequest, nil)
		return
	}
	points -= 1
	if err := store.UpdateUserPoints(request.UserID, points); err != nil {
		respondWithError(w, "Internal server error", http.StatusBadRequest, err)
		return
	}
	if err := store.InsertUserEpisode(request.UserID, request.VideoID, request.Episode); err != nil {
		respondWithError(w, "Internal server error", http.StatusBadRequest, err)
		return
	}
	store.SetCORS(&w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: "Success"})
}

func Episodes(w http.ResponseWriter, r *http.Request) {
	var request EpisodesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, "Invalid request format", http.StatusBadRequest, err)
		return
	}
	if !store.CheckToken(request.Token) {
		respondWithError(w, "Invalid request format", http.StatusBadRequest, nil)
		return
	}
	episodes, err := store.GetUserEpisodeList(request.UserID)
	if err != nil {
		respondWithError(w, "Internal server error", http.StatusBadRequest, err)
		return
	}
	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EpisodesResponse{VideoEpisodeList: episodes})
}

func respondWithError(w http.ResponseWriter, message string, statusCode int, err error) {
	if err != nil {
		log.Println("Episodes:"+message, err)
	} else {
		log.Println("Episodes:" + message)
	}
	http.Error(w, message, statusCode)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	store.SetCORS(&w)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
