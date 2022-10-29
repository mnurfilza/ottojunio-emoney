package transaction

import (
	"e-money-svc/delivery/authentication"
	"e-money-svc/delivery/http"
	"e-money-svc/delivery/route"
	"e-money-svc/repository/client"
	"e-money-svc/repository/db"
	"e-money-svc/usecase/transaction"
)

func TransactionRoute(cll *route.NewRouteGrup) {
	balanceRepo := db.NewBalanceUsers(db.EmoneyDb)
	repoTrans := db.NewDatabaseTransaction(db.EmoneyDb)
	repoBiller := client.NewHttpClient()
	user := db.NewUserClient(db.EmoneyDb)
	uc := transaction.NewTransactionUsecase(repoTrans, repoBiller, user, balanceRepo)
	svc := http.NewTransactionDelivery(uc)
	cll.Transaction.GET("/detail/:trxId", authentication.AutheticateToken(), svc.GetDetailTransaction)
	cll.Transaction.POST("/confirmation", authentication.AutheticateToken(), svc.ConfirmTransaction)
	cll.Transaction.GET("/history", authentication.AutheticateToken(), svc.GetHistoryTransaction)
	cll.Transaction.GET("/inquiry", authentication.AutheticateToken(), svc.GetTransactionsBiller)
}
