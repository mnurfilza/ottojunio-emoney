package model

type BalancesModel struct {
	ID      int    `gorm:"column:id;type:int;autoIncrement"`
	UserId  string `gorm:"column:userId;type:varchar(250)"`
	Balance string `gorm:"column:balance;type:varchar(250)"`
}

func (*BalancesModel) TableName() string {
	return "balance"
}

type TopUpRequest struct {
	Username string `json:"username"`
	Balance  string `json:"balance"`
}

type BalanceInformationResponse struct {
	Balance string `json:"balance"`
}
