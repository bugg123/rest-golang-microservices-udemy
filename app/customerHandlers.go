package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bugg123/rest-golang-microservices-udemy/domain"
	"github.com/bugg123/rest-golang-microservices-udemy/errs"
	"github.com/bugg123/rest-golang-microservices-udemy/service"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("vim-go")
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []domain.Customer
	var err *errs.AppError
	status := r.URL.Query().Get("status")
	if status == "active" {
		customers, err = ch.service.GetCustomerByStatus("1")
	} else if status == "inactive" {
		customers, err = ch.service.GetCustomerByStatus("0")
	} else {
		customers, err = ch.service.GetAllCustomer()
	}
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
