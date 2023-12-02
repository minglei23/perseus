package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token     string
	ID        int
	Email     string
	Activated bool
	VIP       bool
}

func Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Login: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	id, activated, vip, err := store.GetUserIdByEmailAndPassword(request.Email, request.Password)
	if err != nil || id == -1 {
		log.Println("Login: authentication failed:", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token := store.CreateToken(strconv.Itoa(id))
	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token, ID: id, Email: request.Email, Activated: activated, VIP: vip})
}
