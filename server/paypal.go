package server

import (
	"Perseus/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type AmountDetail struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type SaleResource struct {
	CustomID string       `json:"custom_id"`
	Amount   AmountDetail `json:"amount"`
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
	if !store.VerifyPayPal(r.Header) {
		log.Println("Invalid webhook signature")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("Received PayPal webhook notification: %+v\n", notification)
	userID, err := strconv.Atoi(notification.Resource.CustomID)
	if err != nil {
		log.Printf("Error converting userID (%s) to int: %v\n", notification.Resource.CustomID, err)
		http.Error(w, "Error processing metadata", http.StatusInternalServerError)
		return
	}

	log.Printf("CustomID: %d\n", userID)
	log.Printf("Amount: %s\n", notification.Resource.Amount.Value)
	log.Printf("Currency: %s\n", notification.Resource.Amount.CurrencyCode)

	amount, err := strconv.ParseFloat(notification.Resource.Amount.Value, 64)
	if err != nil {
		log.Printf("Error parsing amount: %v\n", err)
		http.Error(w, "Error processing payment amount", http.StatusInternalServerError)
		return
	}

	basicPoints := int64(amount * 50)
	var bonus float64 = 1.1
	if basicPoints >= 1500 {
		bonus = 1.5
	} else if basicPoints >= 1000 {
		bonus = 1.3
	} else if basicPoints >= 250 {
		bonus = 1.2
	}
	additionalPoints := int(float64(basicPoints) * bonus)

	points, err := store.GetPoints(userID)
	if err != nil {
		log.Printf("PayPalWebhook: GetPoints error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	points += additionalPoints
	if err := store.UpdateUserPoints(userID, points); err != nil {
		log.Printf("PayPalWebhook: UpdateUserPoints error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
