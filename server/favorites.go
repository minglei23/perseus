package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type FavoritesRequest struct {
	Token   string
	UserID  int
	VideoID int
}

type FavoritesResponse struct {
	FavoritesList []store.Video
}

func RecordFavorites(w http.ResponseWriter, r *http.Request) {
	var request FavoritesRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("RecordFavorites: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	if !store.CheckToken(request.Token) {
		log.Println("RecordFavorites: check token failed:")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	err = store.InsertUserLike(request.UserID, request.VideoID)
	if err != nil {
		log.Println("RecordFavorites:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	store.SetCORS(&w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: "Success"})
}

func Favorites(w http.ResponseWriter, r *http.Request) {
	var request FavoritesRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Favorites: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	if !store.CheckToken(request.Token) {
		log.Println("Favorites: check token failed:")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	favorites, err := store.GetUserLike(request.UserID)
	if err != nil {
		log.Println("Favorites:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(FavoritesResponse{FavoritesList: favorites})
}
