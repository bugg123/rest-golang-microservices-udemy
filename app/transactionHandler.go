package app

import (
	"encoding/json"
	"net/http"

	"github.com/bugg123/rest-golang-microservices-udemy/dto"
	"github.com/bugg123/rest-golang-microservices-udemy/service"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (h TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	account_id := vars["account_id"]
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error)
	} else {
		request.CustomerId = customer_id
		request.AccountId = account_id
		t, err := h.service.NewTransaction(request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, t)
		}
	}
}
