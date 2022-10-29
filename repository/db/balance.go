package db

import (
	"e-money-svc/domain/model"
	"e-money-svc/domain/repository"
	"gorm.io/gorm"
)

type balanceUsers struct {
	db *gorm.DB
}

func NewBalanceUsers(db *gorm.DB) repository.Balances {
	db.Debug().AutoMigrate(&model.BalancesModel{})
	return &balanceUsers{db: db}
}

func (b balanceUsers) TopUpBalances(in model.BalancesModel) error {
	var req model.BalancesModel
	return b.db.Debug().Model(&req).Where("userId = ?", in.UserId).Update("balance", in.Balance).Error
}

func (b balanceUsers) GetUserBalance(userId string) (*model.BalancesModel, error) {
	var res model.BalancesModel

	if err := b.db.Debug().Where("userId = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (b balanceUsers) InsertBalance(in model.BalancesModel) error {
	return b.db.Debug().Create(&in).Error
}
