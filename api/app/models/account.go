package models

import (
	_ "github.com/jinzhu/gorm"
)

var EnumAccountStatus = struct {
	Enabled string
	Disabled string
	Pending string
	PasswordChange string
	} {
	"ENABLED",
	"DISABLED",
	"PENDING",
	"PASSWORD_CHANGE",
}

// User model
type Account struct {
	Id             string `json:"id"`
	Name           string `gorm:"size:255" json:"name"`
	Email          string `gorm:"type:varchar(100);unique_index" json:"email"`
	Status         string `json:"status"`
}

func (Account) TableName() string {
	return "k_account"
}