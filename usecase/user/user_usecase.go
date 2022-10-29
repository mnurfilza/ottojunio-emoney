package user

import (
	"e-money-svc/domain/model"
	"e-money-svc/domain/repository"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type UserUsecase struct {
	client     repository.Users
	blncClient repository.Balances
}

func NewUsecaseUser(client repository.Users, blncClient repository.Balances) UserUsecaseRepo {
	return &UserUsecase{client: client, blncClient: blncClient}
}

func (u UserUsecase) GetUserInfo(username string) (*model.Users, error) {

	res, err := u.client.GetAccount(username)
	if err != nil {
		return nil, err
	}
	return u.client.GetAccountInfo(strconv.Itoa(res.ID))
}

func (u UserUsecase) UpdateUserInfo(in model.Users) error {
	return u.client.UpdateDataUser(in)
}

func (u UserUsecase) Register(in model.RegisterUser) error {
	//check if username already user
	res, err := u.client.GetAccount(in.Username)
	if err != nil {
		return err
	}

	if res.Username != "" {
		return errors.New("Username Already Use")
	}
	resp, err := u.client.RegisterUser(model.AccountUser{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return err
	}

	//insert info user

	currentDate, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		return nil
	}
	user, err := u.client.InsertUser(model.Users{
		AccountId: strconv.Itoa(resp.ID),
		Name:      in.Name,
		Email:     in.Email,
		BirthDate: currentDate,
	})
	if err != nil {
		return err
	}

	// init balance
	if err := u.blncClient.InsertBalance(model.BalancesModel{
		UserId:  strconv.Itoa(user.ID),
		Balance: "0",
	}); err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) Login(in model.AccountUser) (*model.AccountUser, error) {
	//check password and username
	res, err := u.client.GetAccount(in.Username)
	if err != nil {
		return nil, err
	}

	if res.Username == "" {
		return nil, errors.New("Account unregistered, please register your account")
	}

	return res, nil
}
