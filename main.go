package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/subscriptions/cities/{city}/temperature/average", HandlerGetSubscriptionsCitiesTemperatureAverage)
		r.Get("/subscriptions", HandlerGetSubscriptions)
		r.Put("/subscriptions/cities/{city}", HandlerPutSubscriptionsCities)
		r.Delete("/subscriptions/cities/{city}", HandlerDeleteSubscriptionsCities)
	})

	http.ListenAndServe(":8080", r)
}
