package controllers

import (
	"github.com/revel/revel"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/Zeloid/revelation/app/models"
	"strconv"
)

type Book struct {
	gormc.TxnController
}

func (c Book) Index() revel.Result {
	var books = []models.Book{}
	c.Txn.Find(&books)
	return c.RenderJSON(books)
}

func (c Book) GetOne() revel.Result {
	var book = models.Book{}
	bookId := c.Params.Route.Get("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		return c.RenderJSON(struct{error string}{error: "Unknown Book"})
	}

	c.Txn.Where("id=?", id).Take(&book)
	return c.RenderJSON(book)
}

