package balance

import (
	"e-money-svc/domain/model"
	"e-money-svc/domain/repository"
	"strconv"
	"time"
)

type balanceUsecase struct {
	repo  repository.Balances
	user  repository.Users
	trans repository.TransactionsRepo
}

func NewBalanceUsecase(repo repository.Balances, trans repository.TransactionsRepo, user repository.Users) BalanceUsecaseRepo {
	return &balanceUsecase{repo: repo, trans: trans, user: user}
}

func (b balanceUsecase) GetBalancesUser(username string) (*model.BalanceInformationResponse, error) {
	account, err := b.user.GetAccount(username)
	if err != nil {
		return nil, err
	}

	user, err := b.user.GetAccountInfo(strconv.Itoa(account.ID))
	if err != nil {
		return nil, err
	}

	res, err := b.repo.GetUserBalance(strconv.Itoa(user.ID))
	if err != nil {
		return nil, err
	}
	return &model.BalanceInformationResponse{Balance: res.Balance}, nil
}

func (b balanceUsecase) TopUpBalance(in model.TopUpRequest) error {

	//get userID
	account, err := b.user.GetAccount(in.Username)
	if err != nil {
		return err
	}

	user, err := b.user.GetAccountInfo(strconv.Itoa(account.ID))
	if err != nil {
		return err
	}

	//get last balance
	res, err := b.repo.GetUserBalance(strconv.Itoa(user.ID))
	if err != nil {
		return err
	}

	//format to int
	lastBalance, err := strconv.Atoi(res.Balance)
	if err != nil {
		return err
	}
	topUpBalance, err := strconv.Atoi(in.Balance)
	if err != nil {
		return err
	}

	//count balance
	in.Balance = strconv.Itoa(topUpBalance + lastBalance)

	if err := b.repo.TopUpBalances(model.BalancesModel{
		UserId:  strconv.Itoa(user.ID),
		Balance: in.Balance,
	}); err != nil {
		return err
	}
	//save transaction for history

	timeCurrent, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		return err
	}
	if err := b.trans.InserTransaction(model.Transactions{
		UserId:          strconv.Itoa(user.ID),
		IdBiller:        "",
		Description:     "Top Up Balance",
		DateTransaction: timeCurrent,
		TotalAmount:     in.Balance,
	}); err != nil {
		return err
	}

	return nil

}
