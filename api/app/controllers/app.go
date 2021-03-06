package controllers

import (
	"github.com/revel/revel"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/Zeloid/keeper-chat/api/app/models"
)

type App struct {
	gormc.TxnController
}

func (c App) Index() revel.Result {
	var users = []models.Account{}
	c.Txn.Find(&users)
	return c.RenderJSON(users)
}
