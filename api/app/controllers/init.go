package controllers

import (
	"github.com/revel/revel"
	gorm "github.com/revel/modules/orm/gorm/app"
	"github.com/Zeloid/revelation/app/models"
    "github.com/divan/num2words"
	"fmt"
)

func initializeDB() {
	gorm.DB.AutoMigrate(&models.User{})
	gorm.DB.AutoMigrate(&models.Book{})

	for i := 1; i < 10; i++ {
		model := models.User{Name: fmt.Sprintf("Demo %d", i), Email: fmt.Sprintf("demo%d@demo.com", i)}
		model.SetNewPassword("demo")
		model.Active = true
		gorm.DB.Create(&model)
	}

	for i := 1; i < 10; i++ {
		model := models.Book{
			Title: fmt.Sprintf("%s book", num2words.ConvertAnd(i)),
			Author:fmt.Sprintf("Book %s", num2words.Convert(i)),
			ISBN:uint(i),
		}
		gorm.DB.Create(&model)
	}
}

func init() {
	revel.OnAppStart(initializeDB)
}
