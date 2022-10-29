package route

import "github.com/gin-gonic/gin"

type NewRouteGrup struct {
	User        *gin.RouterGroup
	Transaction *gin.RouterGroup
}

func SetupRoute(router *gin.Engine) *NewRouteGrup {
	return &NewRouteGrup{
		User:        router.Group("emoney/user"),
		Transaction: router.Group("emoney/transaction"),
	}
}
