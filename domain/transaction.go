package domain

import (
	"github.com/bugg123/rest-golang-microservices-udemy/dto"
	"github.com/bugg123/rest-golang-microservices-udemy/errs"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	CustomerId      string
	Amount          float64
	TransactionType string
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *Account, *errs.AppError)
}

func (t Transaction) ToNewTransactionResponseDto(account *Account) dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId: t.TransactionId,
		Amount:        account.Amount,
	}
}
