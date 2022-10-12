package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func HandlerGetSubscriptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello1")
}
func HandlerPutSubscriptionsCities(w http.ResponseWriter, r *http.Request) {
	fmt.Println(chi.URLParam(r, "city"))
}
func HandlerDeleteSubscriptionsCities(w http.ResponseWriter, r *http.Request) {
	fmt.Println(chi.URLParam(r, "city"))
}
func HandlerGetSubscriptionsCitiesTemperatureAverage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(chi.URLParam(r, "city"))
}
