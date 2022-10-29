package model

import "time"

type Transactions struct {
	ID              int       `gorm:"column:id;type:int;autoIncrement"`
	UserId          string    `gorm:"column:userId;type:varchar(250)"`
	IdBiller        string    `gorm:"column:billerId;varchar(250)" json:"idBiller"`
	Description     string    `gorm:"column:description;type:varchar(250)" json:"description"`
	DateTransaction time.Time `gorm:"column:dateTransaction;type:dateTime"`
	TotalAmount     string    `gorm:"column:totalAmount;type:varchar(250)"`
}

func (*Transactions) TableName() string {
	return "transaction"
}

type HistoryTransactionResponse struct {
	UserID string    `json:"userId"`
	Name   string    `json:"name"`
	Biller []*Biller `jsön:"biller"`
}

type DetailTransaction struct {
	UserID     string  `json:"userId"`
	Name       string  `json:"name"`
	Biller     *Biller `json:"biller"`
	Keterangan string  `json:"keterangan,omitempty"`
}

type ConfirmationTransactionReq struct {
	IdBiller string `json:"idBiller"`
	Desc     string `json:"desc"`
	Username string
}
