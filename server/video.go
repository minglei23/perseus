package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type UserVideoRequest struct {
	UserID  int
	VideoID int
	// Code = 1: Get/Record the video that user liked
	// Code = 2: Get/Record the video that user watched
	Code int
}

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

func UserVideo(w http.ResponseWriter, r *http.Request) {
	var userVideoRequest UserVideoRequest
	if err := json.NewDecoder(r.Body).Decode(&userVideoRequest); err != nil {
		log.Println("UserVideo: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("perseus")
	if err != nil {
		log.Println("UserVideo: get cookie error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	checkCookie, err := store.CheckCookie(*cookie)
	if !checkCookie || err != nil {
		log.Println("UserVideo: check cookie failed:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if userVideoRequest.Code == 1 {
		err = store.InsertUserLike(userVideoRequest.UserID, int(userVideoRequest.VideoID))
	} else {
		err = store.InsertUserHistory(userVideoRequest.UserID, int(userVideoRequest.VideoID))
	}
	if err != nil {
		log.Println("UserVideo:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	store.SetCORS(&w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: "Success"})
}

func UserVideoList(w http.ResponseWriter, r *http.Request) {
	var userVideoRequest UserVideoRequest
	if err := json.NewDecoder(r.Body).Decode(&userVideoRequest); err != nil {
		log.Println("UserVideoList: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("perseus")
	if err != nil {
		log.Println("UserVideoList: get cookie error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	checkCookie, err := store.CheckCookie(*cookie)
	if !checkCookie || err != nil {
		log.Println("UserVideoList: check cookie failed:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var userVideoList []store.Video
	if userVideoRequest.Code == 1 {
		userVideoList, err = store.GetUserLike(userVideoRequest.UserID)
	} else {
		userVideoList, err = store.GetUserHistoryLstMonth(userVideoRequest.UserID)
	}
	if err != nil {
		log.Println("UserVideoList:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VideoListResponse{VideoList: userVideoList})
}
