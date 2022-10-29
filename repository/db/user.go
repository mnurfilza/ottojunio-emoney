package db

import (
	"e-money-svc/domain/model"
	"e-money-svc/domain/repository"
	"errors"
	"gorm.io/gorm"
)

type usersClientDatabase struct {
	db *gorm.DB
}

func NewUserClient(db *gorm.DB) repository.Users {
	db.Debug().AutoMigrate(&model.Users{}, &model.AccountUser{})
	return &usersClientDatabase{db: db}
}

func (u usersClientDatabase) RegisterUser(in model.AccountUser) (*model.AccountUser, error) {
	if err := u.db.Debug().Create(&in).Error; err != nil {
		return nil, err
	}
	return &in, nil
}

func (u usersClientDatabase) UpdateDataUser(in model.Users) error {
	var user model.Users
	return u.db.Debug().Model(&user).Where("id = ?", in.ID).Updates(in).Error
}

func (u usersClientDatabase) GetAccountInfo(id string) (*model.Users, error) {
	var res model.Users
	if err := u.db.Debug().Where("accountId = ? OR id = ?", id, id).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (u usersClientDatabase) GetAccount(username string) (*model.AccountUser, error) {
	var res model.AccountUser
	if err := u.db.Debug().Where("username = ?", username).Find(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AccountUser{}, nil
		}

		return nil, err
	}

	return &res, nil
}

func (u usersClientDatabase) InsertUser(in model.Users) (*model.Users, error) {
	if err := u.db.Debug().Create(&in).Error; err != nil {
		return nil, err
	}
	return &in, nil
}
