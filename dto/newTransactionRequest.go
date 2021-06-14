package dto

import (
	"strings"

	"github.com/bugg123/rest-golang-microservices-udemy/errs"
)

type Transaction struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (r Transaction) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidateError("Transactions must be postive")
	}
	if strings.ToLower(r.TransactionType) != "withdrawal" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidateError("Account type must be withdrawal or deposit")
	}
	return nil
}
