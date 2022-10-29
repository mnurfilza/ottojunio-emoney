package transaction

import (
	"e-money-svc/domain/model"
	"e-money-svc/domain/repository"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type transactionUsecase struct {
	repo    repository.TransactionsRepo
	biller  repository.BillerServiceClient
	user    repository.Users
	balance repository.Balances
}

func NewTransactionUsecase(repo repository.TransactionsRepo, biller repository.BillerServiceClient, user repository.Users, balance repository.Balances) TransactionUsecaseRepo {
	return &transactionUsecase{repo: repo, biller: biller, user: user, balance: balance}
}

func (t transactionUsecase) InsertTransaction(in model.ConfirmationTransactionReq) error {
	var err error
	dateTrx, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		return err
	}

	account, err := t.user.GetAccount(in.Username)
	if err != nil {
		return err
	}

	user, err := t.user.GetAccountInfo(strconv.Itoa(account.ID))
	if err != nil {
		return err
	}
	//get data biller
	res, err := t.biller.GetDetailBiller(in.IdBiller)
	if err != nil {
		return err
	}

	if err := t.repo.InserTransaction(model.Transactions{
		UserId:          strconv.Itoa(user.ID),
		IdBiller:        strconv.Itoa(res.Data.ID),
		Description:     res.Data.Desc,
		DateTransaction: dateTrx,
		TotalAmount:     strconv.Itoa(res.Data.Price + res.Data.Fee),
	}); err != nil {
		return err
	}

	//decrease balance
	blnc, err := t.balance.GetUserBalance(strconv.Itoa(user.ID))
	if err != nil {
		return err
	}

	currentBlance, err := strconv.Atoi(blnc.Balance)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		return err
	}

	blnc.Balance = strconv.Itoa(currentBlance - (res.Data.Price + res.Data.Fee))

	if err := t.balance.TopUpBalances(*blnc); err != nil {
		return err
	}

	return nil
}

func (t transactionUsecase) GetHistoryTransaction(username, startDate, endDate string) ([]model.Transactions, error) {
	account, err := t.user.GetAccount(username)
	if err != nil {
		return nil, err
	}

	user, err := t.user.GetAccountInfo(strconv.Itoa(account.ID))
	if err != nil {
		return nil, err
	}
	return t.repo.TransactionHistory(strconv.Itoa(user.ID), startDate, endDate)
}

func (t transactionUsecase) GetDetailTransaction(id string) (*model.DetailTransaction, error) {
	res, err := t.repo.GetDetailTransaction(id)
	if err != nil {
		return nil, err
	}

	user, err := t.user.GetAccountInfo(res.UserId)
	if err != nil {
		return nil, err
	}

	bilRes, err := t.biller.GetDetailBiller(res.IdBiller)
	if err != nil {
		return nil, err
	}

	return &model.DetailTransaction{
		UserID: res.UserId,
		Name:   user.Name,
		Biller: bilRes.Data,
	}, nil
}

func (t transactionUsecase) GetListTransaction() (map[string][]*model.Biller, error) {
	res := make(map[string][]*model.Biller)
	resp, err := t.biller.GetListBiller()
	if err != nil {
		return nil, err
	}
	var bill []*model.Biller
	for len(resp.Data) != 0 {
		_, ok := res[strings.ToLower(resp.Data[0].Category)]
		if !ok || ok {
			if len(bill) > 0 {
				if bill[len(bill)-1].Category != resp.Data[0].Category {
					bill = nil
				}
			}
			bill = append(bill, resp.Data[0])
			res[strings.ToLower(resp.Data[0].Category)] = bill

		}

		resp.Data = resp.Data[1:]
	}
	return res, nil

}
