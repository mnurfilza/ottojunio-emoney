package http

import (
	"e-money-svc/domain/model"
	"e-money-svc/usecase/balance"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BalanceDelivery struct {
	blnc balance.BalanceUsecaseRepo
}

func NewBalanceDelivery(blnc balance.BalanceUsecaseRepo) *BalanceDelivery {
	return &BalanceDelivery{blnc: blnc}
}

func (b *BalanceDelivery) TopUpBalance(c *gin.Context) {
	var req model.TopUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not Allowed Access "})
		return
	}

	req.Username = ctx.(string)
	if err := b.blnc.TopUpBalance(req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Top Success"})
	return
}

func (b *BalanceDelivery) GetBalanceInfo(c *gin.Context) {
	ctx, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not Allowed Access "})
		return
	}
	res, err := b.blnc.GetBalancesUser(ctx.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return

}
