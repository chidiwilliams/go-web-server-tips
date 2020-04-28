package repository

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/chidiwilliams/go-web-server-tips/errors"
	"github.com/chidiwilliams/go-web-server-tips/models"
)

var (
	errBookNotFound = errors.Error("book not found")
)

type Repository interface {
	GetBook(id bson.ObjectId) (*models.Book, error)
	CreateBook(book models.Book) error
}
