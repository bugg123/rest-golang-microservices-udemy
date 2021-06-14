package service

import (
	"github.com/bugg123/rest-golang-microservices-udemy/domain"
	"github.com/bugg123/rest-golang-microservices-udemy/dto"
	"github.com/bugg123/rest-golang-microservices-udemy/errs"
)

type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	t := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		CustomerId:      req.CustomerId,
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
	}
	newTransaction, err := s.repo.Save(t)
	if err != nil {
		return nil, err
	}
	resp := newTransaction.ToNewTransactionResponseDto()
	return &resp, nil
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
