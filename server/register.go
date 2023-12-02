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
	Code      int
	Token     string
	ID        int
	Email     string
	Activated bool
	VIP       bool
}

func Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Register: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	emailExists, err := store.EmailExist(request.Email)
	if err != nil {
		log.Println("Register: email exist error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if emailExists {
		respondWithCode(w, 2, -1, request.Email)
		return
	}

	id, err := store.InsertUser(request.Password, request.Email)
	if err != nil {
		log.Println("Register: insert user error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithCode(w, 1, id, request.Email)
}

func respondWithCode(w http.ResponseWriter, code int, id int, email string) {
	token := store.CreateToken(strconv.Itoa(id))
	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RegisterResponse{Code: code, Token: token, ID: id, Email: email, Activated: false, VIP: false})
}
