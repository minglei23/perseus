package main

import (
	"Perseus/server"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v76"
)

func main() {

	// Stripe test secret API key.
	stripe.Key = "sk_test_51OFXw4Lvs8YNyX8smm0SPVJwTaxKg31XcafaF8zTPHzzhuDaGOkjvi374NnUdiHf7JMYowaVQ4nK1Y6ac8rhVsTl00UTn0o3Ty"
	http.HandleFunc("/create-checkout-session", server.CreateCheckoutSession)

	http.HandleFunc("/login", server.Login)
	http.HandleFunc("/register", server.Register)

	http.HandleFunc("/reset-email", server.ResetEmail)
	http.HandleFunc("/reset", server.Reset)

	http.HandleFunc("/verify-email", server.VerifyEmail)
	http.HandleFunc("/verify", server.Verify)

	http.HandleFunc("/video-list", server.VideoList)

	http.HandleFunc("/record-favorites", server.RecordFavorites)
	http.HandleFunc("/remove-favorites", server.RemoveFavorites)
	http.HandleFunc("/favorites", server.Favorites)

	http.HandleFunc("/record-history", server.RecordHistory)
	http.HandleFunc("/history", server.History)

	http.HandleFunc("/points", server.Points)
	http.HandleFunc("/already-checkin", server.AlreadyCheckin)
	http.HandleFunc("/checkin", server.Checkin)

	http.HandleFunc("/upload-video-info", server.UploadVideoInfo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
