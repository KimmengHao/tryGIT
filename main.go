package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "token:Hello, New world!")
}

// Initiate web server
func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedOrigins([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedOrigins([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/auth", Authorize)
	log.Fatal(http.ListenAndServe(":9090", handlers.CORS(headers, methods, origins)(router)))
}
