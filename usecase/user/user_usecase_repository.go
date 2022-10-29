package user

import "e-money-svc/domain/model"

type UserUsecaseRepo interface {
	GetUserInfo(userId string) (*model.Users, error)
	UpdateUserInfo(in model.Users) error
	Register(in model.RegisterUser) error
	Login(in model.AccountUser) (*model.AccountUser, error)
}
