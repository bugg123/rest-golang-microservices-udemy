package service

import (
	"github.com/bugg123/rest-golang-microservices-udemy/domain"
	"github.com/bugg123/rest-golang-microservices-udemy/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return d.repo.FindAll()
}

func (d DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return d.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
