package app

import (
	"log"
	"net/http"
	"os"

	"github.com/bugg123/rest-golang-microservices-udemy/domain"
	"github.com/bugg123/rest-golang-microservices-udemy/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	envVars := []string{"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	undefined := make([]string, 0)
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			undefined = append(undefined, envVar)
		}
	}
	if len(undefined) != 0 {
		log.Fatalf("Necessary envVars not defined: %v", undefined)
	}
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}
