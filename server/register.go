package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type RegisterRequest struct {
	Email    string
	Password string
}

type RegisterResponse struct {
	// Code = 1: register successfully
	// Code = 2: email is existing in database
	// Code = 3: register successfully, but login failed
	Code int
}

func Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		log.Println("Register: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Check if email exists
	emailExists, err := store.EmailExist(registerRequest.Email)
	if err != nil {
		log.Println("Register: email exist error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if emailExists {
		respondWithCode(w, 2)
		return
	}

	// Inser user
	id, err := store.InsertUser(registerRequest.Password, registerRequest.Email)
	if err != nil {
		log.Println("Register: insert user error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create cookie
	cookie, err := store.CreateCookie(strconv.FormatInt(id, 10))
	if err != nil {
		respondWithCode(w, 3)
		return
	}

	http.SetCookie(w, &cookie)
	respondWithCode(w, 1)
}

func respondWithCode(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RegisterResponse{Code: code})
}
