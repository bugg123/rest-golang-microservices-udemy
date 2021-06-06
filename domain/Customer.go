package domain

import "github.com/bugg123/rest-golang-microservices-udemy/errs"

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByStatus(string) ([]Customer, *errs.AppError)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Corey", "Morgantown", "12345", "2000-01-01", "1"},
		{"1002", "Caleb", "Portland", "12345", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
