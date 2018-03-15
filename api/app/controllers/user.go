package controllers

import (
	"github.com/revel/revel"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/Zeloid/keeper-chat/api/app/models"
	"net/http"
)

type User struct {
	gormc.TxnController
}

func (c User) Index() revel.Result {
	var users []models.Account
	c.Txn.Find(&users)
	return c.RenderJSON(users)
}

func (c User) Create(account *models.Account) revel.Result {
	c.Validation.Required(account.Name)
	c.Validation.MinSize(account.Name,4)
	c.Validation.MaxSize(account.Name,45)
	c.Validation.Required(account.Email)
	c.Validation.MaxSize(account.Name,45)
	c.Validation.Email(account.Email)

	if c.Validation.HasErrors() {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(c.Validation.Errors)
	}

	account.Id = ""
	account.Status = models.EnumAccountStatus.Pending

	c.Txn.Create(&account)

	return c.RenderJSON(account)
}