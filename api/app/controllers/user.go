package controllers

import (
	"github.com/revel/revel"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/Zeloid/revelation/app/models"
)

type User struct {
	gormc.TxnController
}

func (c User) Index() revel.Result {
	var users = []models.User{}
	c.Txn.Find(&users)
	return c.RenderJSON(users)
}
