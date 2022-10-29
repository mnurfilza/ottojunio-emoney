package http

import (
	"e-money-svc/delivery/authentication"
	"e-money-svc/domain/model"
	"e-money-svc/usecase/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserDelivery struct {
	user user.UserUsecaseRepo
}

func NewUserDelivery(user user.UserUsecaseRepo) *UserDelivery {
	return &UserDelivery{user: user}
}

func (u *UserDelivery) Register(c *gin.Context) {
	var req model.RegisterUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := u.user.Register(req); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.PureJSON(200, gin.H{"message": "Register Success"})
}

func (u *UserDelivery) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res, err := u.user.Login(model.AccountUser{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	token, err := authentication.CreateToken(res.Username)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("x-token", token)
	c.JSON(http.StatusOK, gin.H{"message": "Login Success"})
}

func (u *UserDelivery) GetAccountInfo(c *gin.Context) {
	ctx, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not Allowed Access "})
		return
	}
	res, err := u.user.GetUserInfo(ctx.(string))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)

}
