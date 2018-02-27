package models

import (
	"github.com/jinzhu/gorm"
)

// Book model
type Book struct {
	gorm.Model
	Title          string `gorm:"size:255"`
	Author         string `gorm:"size:255"`
	ISBN           uint   `gorm:"unique_index"`
}

func (Book) TableName() string {
	return "md_book"
}