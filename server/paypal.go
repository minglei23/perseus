package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
)

type SaleResource struct {
	CustomID int `json:"custom_id"`
}

type PayPalWebhookEvent struct {
	ID           string       `json:"id"`
	CreateTime   string       `json:"create_time"`
	ResourceType string       `json:"resource_type"`
	EventType    string       `json:"event_type"`
	Summary      string       `json:"summary"`
	Resource     SaleResource `json:"resource"`
}

func PayPalWebhook(w http.ResponseWriter, r *http.Request) {
	var notification PayPalWebhookEvent
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		log.Printf("Error decoding webhook request: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if store.VerifyPayPal(r.Header) {
		log.Println("Invalid webhook signature")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := notification.Resource.CustomID
	points, err := store.GetPoints(userID)
	if err != nil {
		log.Printf("PayPalWebhook: GetPoints error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	points += 120
	if err := store.UpdateUserPoints(userID, points); err != nil {
		log.Printf("PayPalWebhook: UpdateUserPoints error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
