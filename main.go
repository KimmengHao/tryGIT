package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// EDIT THIS
var (
	BaseURL      = "http://localhost:9090"
	ClientID     = "BhWus13WIhI4HI7loycM42"
	ClientSecret = "HJHgKSrqUuhIepCzNkCx7E82RSTN1m47dqPoS1Lf6VA"
)

type numsResponseData struct {
	UserID string  `json:"UserID"`
	Point  float64 `json:"Point"`
	// Jsondata string `json:"Jsondata"`
}

func Authorize(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "token:Authorize")
}

func Callback(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "token:callabcak")
}

func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedOrigins([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedOrigins([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/auth", Authorize)
	router.HandleFunc("/callback", Callback)
	log.Fatal(http.ListenAndServe(":9090", handlers.CORS(headers, methods, origins)(router)))
}
