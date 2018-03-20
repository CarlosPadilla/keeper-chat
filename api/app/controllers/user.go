package controllers

import (
	"github.com/revel/revel"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/Zeloid/keeper-chat/api/app/models"
	"net/http"
	"github.com/Zeloid/keeper-chat/api/app/components"
)

type User struct {
	gormc.TxnController
}

func (c User) Index() revel.Result {
	var users []models.Account
	c.Txn.Find(&users)
	return c.RenderJSON(users)
}


/**
 * Respond to the post to the account endpoint to create a new account
 * Unmarshal the JSON body as an models.account struct
 */
func (c User) Create(account *models.Account) revel.Result {
	// validations
	c.Validation.Required(account.Name)
	c.Validation.MinSize(account.Name,4)
	c.Validation.MaxSize(account.Name,45)
	c.Validation.Required(account.Email)
	c.Validation.MaxSize(account.Name,45)
	c.Validation.Email(account.Email)

	if c.Validation.HasErrors() {
		c.Response.Status = http.StatusBadRequest
		errorResponse := components.NewValidationErrorResponse(c.Validation.Errors)

		return c.RenderJSON(errorResponse)
	}

	// fields with fixed values on creation
	account.Id = ""
	account.Status = models.EnumAccountStatus.Pending

	// make the insert
	c.Txn.Create(&account)

	// return model as JSON
	return c.RenderJSON(account)
}