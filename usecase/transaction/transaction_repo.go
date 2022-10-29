package transaction

import (
	"e-money-svc/domain/model"
)

type TransactionUsecaseRepo interface {
	InsertTransaction(in model.ConfirmationTransactionReq) error
	GetHistoryTransaction(userId, startDate, endDate string) ([]model.Transactions, error)
	GetDetailTransaction(trxId string) (*model.DetailTransaction, error)
	GetListTransaction() (map[string][]*model.Biller, error)
}
