package repository

import (
	"e-money-svc/domain/model"
)

type TransactionsRepo interface {
	TransactionHistory(userId string, startDate, endDate string) ([]model.Transactions, error)
	InserTransaction(in model.Transactions) error
	GetDetailTransaction(userId string) (*model.Transactions, error)
}
