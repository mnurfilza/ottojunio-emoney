package repository

import "e-money-svc/domain/model"

type Balances interface {
	//updaet to database
	TopUpBalances(in model.BalancesModel) error
	GetUserBalance(userId string) (*model.BalancesModel, error)
	InsertBalance(in model.BalancesModel) error
}
