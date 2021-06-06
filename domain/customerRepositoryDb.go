package domain

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/bugg123/rest-golang-microservices-udemy/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	return c.getCustomersByQuery(findAllSql)
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

func (d CustomerRepositoryDb) ByStatus(status string) ([]Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	return d.getCustomersByQuery(customerSql, status)
}

func (d CustomerRepositoryDb) getCustomersByQuery(sqlQuery string, args ...interface{}) ([]Customer, *errs.AppError) {
	rows, err := d.client.Query(sqlQuery, args...)
	if err != nil {
		log.Printf("Unable to query customer table: %v", err)
		return nil, errs.NewAppError("Unable to query customer table", http.StatusInternalServerError)
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			return nil, errs.NewAppError("Error while scanning customers", http.StatusInternalServerError)
		}
		customers = append(customers, c)
	}
	return customers, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	db, err := sql.Open("mysql", "root:codecamp@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDb{db}
}
