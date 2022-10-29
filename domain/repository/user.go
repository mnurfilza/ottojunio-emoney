package repository

import "e-money-svc/domain/model"

type Users interface {
	RegisterUser(in model.AccountUser) (*model.AccountUser, error)
	UpdateDataUser(in model.Users) error
	GetAccountInfo(id string) (*model.Users, error)
	GetAccount(username string) (*model.AccountUser, error)
	InsertUser(in model.Users) (*model.Users, error)
}
