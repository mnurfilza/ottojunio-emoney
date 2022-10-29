package user

import (
	"e-money-svc/delivery/authentication"
	"e-money-svc/delivery/http"
	"e-money-svc/delivery/route"
	"e-money-svc/repository/db"
	"e-money-svc/usecase/balance"
	user2 "e-money-svc/usecase/user"
)

func RouterUserGroup(colection *route.NewRouteGrup) {
	userRepo := db.NewUserClient(db.EmoneyDb)
	balanceRepo := db.NewBalanceUsers(db.EmoneyDb)
	transRepo := db.NewDatabaseTransaction(db.EmoneyDb)
	usrUc := user2.NewUsecaseUser(userRepo, balanceRepo)
	balanceUc := balance.NewBalanceUsecase(balanceRepo, transRepo, userRepo)

	userDlv := http.NewUserDelivery(usrUc)
	balanceSvc := http.NewBalanceDelivery(balanceUc)

	colection.User.POST("/register", userDlv.Register)
	colection.User.POST("/login", userDlv.Login)
	colection.User.GET("/detail", authentication.AutheticateToken(), userDlv.GetAccountInfo)

	//topup and get balance

	colection.User.GET("/balance", authentication.AutheticateToken(), balanceSvc.GetBalanceInfo)
	colection.User.POST("/topUp", authentication.AutheticateToken(), balanceSvc.TopUpBalance)
}
