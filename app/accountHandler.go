package app

import (
	"encoding/json"
	"net/http"

	"github.com/bugg123/rest-golang-microservices-udemy/dto"
	"github.com/bugg123/rest-golang-microservices-udemy/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error)
	} else {
		request.CustomerId = customer_id
		account, err := h.service.NewAccount(request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
