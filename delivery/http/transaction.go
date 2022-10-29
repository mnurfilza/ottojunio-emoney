package http

import (
	"e-money-svc/domain/model"
	"e-money-svc/usecase/transaction"
	ggin "github.com/gin-gonic/gin"
	"net/http"
)

type TransactionDelivery struct {
	uc transaction.TransactionUsecaseRepo
}

func NewTransactionDelivery(uc transaction.TransactionUsecaseRepo) *TransactionDelivery {
	return &TransactionDelivery{uc: uc}
}

func (s *TransactionDelivery) ConfirmTransaction(c *ggin.Context) {
	var req model.ConfirmationTransactionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ggin.H{"message": err.Error()})
		return
	}
	ctx, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ggin.H{"message": "Not Allowed Access "})
		return
	}

	req.Username = ctx.(string)

	if err := s.uc.InsertTransaction(req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ggin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ggin.H{"message": "Success Transcation"})
	return
}

func (s *TransactionDelivery) GetDetailTransaction(c *ggin.Context) {
	trxId := c.Param("trxId")

	res, err := s.uc.GetDetailTransaction(trxId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ggin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func (s *TransactionDelivery) GetHistoryTransaction(c *ggin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	ctx, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ggin.H{"message": "Not Allowed Access "})
		return
	}
	res, err := s.uc.GetHistoryTransaction(ctx.(string), startDate, endDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ggin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (s *TransactionDelivery) GetTransactionsBiller(c *ggin.Context) {
	res, err := s.uc.GetListTransaction()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
