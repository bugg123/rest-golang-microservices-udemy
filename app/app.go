package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()
	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
