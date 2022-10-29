package db

import (
	"e-money-svc/domain/model"
	"e-money-svc/domain/repository"
	"gorm.io/gorm"
)

type transactionClientDB struct {
	db *gorm.DB
}

func NewDatabaseTransaction(db *gorm.DB) repository.TransactionsRepo {
	db.Debug().AutoMigrate(&model.Transactions{})
	return &transactionClientDB{db: db}
}

func (t transactionClientDB) TransactionHistory(userId string, startDate, endDate string) ([]model.Transactions, error) {
	var res []model.Transactions
	if err := t.db.Debug().Where("userId = ? AND (dateTransaction BETWEEN ? AND ?)", userId, startDate, endDate).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (t transactionClientDB) InserTransaction(in model.Transactions) error {
	return t.db.Debug().Create(&in).Error
}

func (t transactionClientDB) GetDetailTransaction(id string) (*model.Transactions, error) {
	var res model.Transactions
	if err := t.db.Debug().Where("id = ?", id).Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
