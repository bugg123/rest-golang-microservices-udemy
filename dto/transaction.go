package dto

import (
	"strings"

	"github.com/bugg123/rest-golang-microservices-udemy/errs"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.IsTransactionTypeWithdrawal() && r.IsTransactionTypeDeposit() {
		return errs.NewValidateError("Account type must be withdrawal or deposit")
	}
	if r.Amount < 0 {
		return errs.NewValidateError("Transactions must be postive")
	}
	return nil
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return strings.ToLower(r.TransactionType) == "withdrawal"
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return strings.ToLower(r.TransactionType) == "deposit"
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
