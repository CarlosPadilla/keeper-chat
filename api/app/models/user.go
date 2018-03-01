package models

import (
	_ "github.com/jinzhu/gorm"
)

// User model
type User struct {
	Id             string
	Name           string `gorm:"size:255"`
	Email          string `gorm:"type:varchar(100);unique_index"`
	Status         string
}

func (User) TableName() string {
	return "k_account"
}