package service

import (
	"github.com/bugg123/rest-golang-microservices-udemy/domain"
	"github.com/bugg123/rest-golang-microservices-udemy/dto"
	"github.com/bugg123/rest-golang-microservices-udemy/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	GetCustomerByStatus(string) ([]dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := d.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var resp []dto.CustomerResponse
	for _, customer := range customers {
		resp = append(resp, customer.ToDto())
	}
	return resp, nil
}

func (d DefaultCustomerService) GetCustomerByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := d.repo.ByStatus(status)
	if err != nil {
		return nil, err
	}
	var resp []dto.CustomerResponse
	for _, customer := range customers {
		resp = append(resp, customer.ToDto())
	}
	return resp, nil
}

func (d DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := d.repo.ById(id)
	if err != nil {
		return nil, err
	}
	resp := c.ToDto()
	return &resp, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
