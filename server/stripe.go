package server

import (
	"Perseus/store"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/webhook"
)

type StripeRequest struct {
	ID         int
	Amount     int64
	ProductID  string
	SuccessURL string
	CancelURL  string
}

type StripeResponse struct {
	SessionId string
}

func CreateStripePayment(w http.ResponseWriter, r *http.Request) {
	var request StripeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("CreateStripePayment: json decoder error:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(request.ProductID),
				Quantity: stripe.Int64(request.Amount),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(request.SuccessURL),
		CancelURL:  stripe.String(request.CancelURL),
		Metadata: map[string]string{
			"userID": strconv.Itoa(request.ID),
		},
	}
	s, err := session.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.SetCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(StripeResponse{SessionId: s.ID})
}

func StripeWebhook(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), store.StripeWebhook)
	if err != nil {
		log.Printf("Error verifying webhook signature: %v\n", err)
		http.Error(w, "Error verifying webhook signature", http.StatusBadRequest)
		return
	}

	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			log.Printf("Error unmarshalling session: %v\n", err)
			http.Error(w, "Error processing checkout session", http.StatusInternalServerError)
			return
		}
		log.Panicln("Stripe Webhook:", session)

		userIDStr, ok := session.Metadata["userID"]
		if !ok {
			log.Println("Error: userID not found in metadata")
			http.Error(w, "userID not found in metadata", http.StatusInternalServerError)
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			log.Printf("Error converting userID (%s) to int: %v\n", userIDStr, err)
			http.Error(w, "Error processing metadata", http.StatusInternalServerError)
			return
		}

		additionalPoints := int(session.AmountTotal / 2)
		var bonus float64 = 1.1
		if additionalPoints >= 1500 {
			bonus = 1.5
		} else if additionalPoints >= 1000 {
			bonus = 1.3
		} else if additionalPoints >= 250 {
			bonus = 1.2
		}
		additionalPoints = int(float64(additionalPoints) * bonus)
		log.Printf("User %s buy points %d\n", userIDStr, additionalPoints)

		points, err := store.GetPoints(userID)
		if err != nil {
			log.Printf("StripeWebhook: GetPoints error: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		points += additionalPoints
		if err := store.UpdateUserPoints(userID, points); err != nil {
			log.Printf("StripeWebhook: UpdateUserPoints error: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
