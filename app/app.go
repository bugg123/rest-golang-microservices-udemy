package app

import (
	"log"
	"net/http"

	"github.com/bugg123/rest-golang-microservices-udemy/domain"
	"github.com/bugg123/rest-golang-microservices-udemy/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
