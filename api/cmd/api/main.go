package main

import (
	"fmt"
	"log"
	"net/http"

	"SettleGuard/internal/api"
)

func main() {
	// Define the checkout route
	http.HandleFunc("/api/checkout", enableCORS(api.HandleCheckout))

	port := ":8080"
	fmt.Printf("🚀 (to the moon) SettleGuard API Server running on http://localhost%s\n", port)
	fmt.Printf("🛒 (need a quarter?) Checkout endpoint ready for demo site: http://localhost%s/api/checkout\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// enableCORS Why'd you have to go and make things so complicated?
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust for production
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
