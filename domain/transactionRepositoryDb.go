package domain

import (
	"strconv"

	"github.com/bugg123/rest-golang-microservices-udemy/errs"
	"github.com/bugg123/rest-golang-microservices-udemy/logger"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDb) Save(t Transaction) (*Transaction, *errs.AppError) {
	sqlInsert := "INSERT INTO transactions (account_id, customer_id, amount, transaction_type) values (?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, t.AccountId, t.CustomerId, t.Amount, t.TransactionType)
	if err != nil {
		logger.Error("Error while creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error getting last id for insert for new transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	t.TransactionId = strconv.FormatInt(id, 10)
	return &t, nil
}

func (d TransactionRepositoryDb) updateAccount(t Transaction) (*Account, *errs.AppError) {
	sqlSelect := "SELECT account FROM accounts where account_id = ?"
	var acct Account
	err := d.client.Select(acct, sqlSelect, t.AccountId)
	if err != nil {
		logger.Error("Unable to select account for transaction")
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}
