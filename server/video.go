package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type VideoListResponse struct {
	VideoList []store.Video
}

func VideoList(w http.ResponseWriter, r *http.Request) {

	videoList, err := store.GetVideoList()
	if err != nil {
		log.Println("videoList: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VideoListResponse{VideoList: videoList})
}
