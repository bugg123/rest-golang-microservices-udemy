package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/bugg123/rest-golang-microservices-udemy/errs"
	"github.com/bugg123/rest-golang-microservices-udemy/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	return c.getCustomersByQuery(findAllSql)
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

func (d CustomerRepositoryDb) ByStatus(status string) ([]Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	return d.getCustomersByQuery(customerSql, status)
}

func (d CustomerRepositoryDb) getCustomersByQuery(sqlQuery string, args ...interface{}) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	err := d.client.Select(&customers, sqlQuery, args...)
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to query customer table: %v", err))
		return nil, errs.NewUnexpectedError("Unable to query customer table")
	}

	return customers, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDb{db}
}
