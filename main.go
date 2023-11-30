package main

import (
	"Perseus/server"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/login", server.Login)
	http.HandleFunc("/register", server.Register)
	http.HandleFunc("/reset-email", server.ResetEmail)
	http.HandleFunc("/reset", server.Reset)
	http.HandleFunc("/verify-email", server.VerifyEmail)
	http.HandleFunc("/verify", server.Verify)

	http.HandleFunc("/video-list", server.VideoList)
	http.HandleFunc("/user-video", server.UserVideo)
	http.HandleFunc("/user-video-list", server.UserVideoList)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
