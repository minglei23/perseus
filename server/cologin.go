package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CoLoginRequest struct {
	Id    string
	Token string
	Type  int
	Email string
}

func CoLogin(w http.ResponseWriter, r *http.Request) {
	var request CoLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("CoLogin: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	id, activated, vip, err := store.GetUserIdByIdAndType(request.Id, request.Type)
	if err != nil {
		log.Println("CoLogin: authentication failed:", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if id == -1 {
		id, err = store.InsertCoUser(request.Id, request.Email, request.Type)
		if err != nil {
			log.Println("CoLogin: insert user error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	token := store.CreateToken(strconv.Itoa(id))
	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token, ID: id, Email: request.Email, Activated: activated, VIP: vip})
}
