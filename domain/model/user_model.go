package model

import "time"

type Users struct {
	ID          int       `gorm:"column:id;type:int;autoIncrement"`
	AccountId   string    `gorm:"column:accountId;type:varchar(255)"`
	Name        string    `gorm:"column:name;type:varchar(250)" json:"name"`
	Address     string    `gorm:"column:address;type:varchar(250)" json:"address"`
	Email       string    `gorm:"column:email;type:varchar(250)" json:"email"`
	PhoneNumber string    `gorm:"column:phoneNumber;type:varchar(255)" json:"phoneNumber"`
	BirthDate   time.Time `gorm:"column:birthDate;type:dateTime" json:"birthDate"`
}

func (*Users) TableName() string {
	return "user"
}

type AccountUser struct {
	ID       int    `gorm:"column:id;type:int;autoIncrement"`
	Username string `gorm:"column:username;type:varchar(255)"`
	Password string `gorm:"column:password;type:varchar(255)"`
}

func (*AccountUser) TableName() string {
	return "account_user"
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token   string
	Message string
}

type RegisterUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
