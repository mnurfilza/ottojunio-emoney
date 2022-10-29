package main

import (
	"e-money-svc/delivery/route"
	"e-money-svc/delivery/route/transaction"
	"e-money-svc/delivery/route/user"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Recovery())
	route := route.SetupRoute(router)
	user.RouterUserGroup(route)
	transaction.TransactionRoute(route)
	router.Run(":8801")
}
