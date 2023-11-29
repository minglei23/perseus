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
	ID        int64
	Email     string
	Activated bool
	VIP       bool
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Login: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	id, activated, vip, err := store.GetUserIdByEmailAndPassword(loginRequest.Email, loginRequest.Password)
	if err != nil || id == -1 {
		log.Println("Login: authentication failed:", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	cookie, err := store.CreateCookie(strconv.FormatInt(id, 10))
	if err != nil {
		log.Println("Login: create cookie error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &cookie)

	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{ID: id, Email: loginRequest.Email, Activated: activated, VIP: vip})
}
