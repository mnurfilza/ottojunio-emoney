package balance

import "e-money-svc/domain/model"

type BalanceUsecaseRepo interface {
	GetBalancesUser(username string) (*model.BalanceInformationResponse, error)
	TopUpBalance(in model.TopUpRequest) error
}
