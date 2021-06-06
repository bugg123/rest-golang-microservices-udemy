package service

import "github.com/bugg123/rest-golang-microservices-udemy/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return d.repo.FindAll()
}

func (d DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return d.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
