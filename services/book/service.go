package book

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/chidiwilliams/go-web-server-tips/models"
	"github.com/chidiwilliams/go-web-server-tips/services/book/repository"
)

type Service interface {
	GetBook(id bson.ObjectId) (*models.Book, error)
	CreateBook(title string) (*models.Book, error)
}

func NewService(repository repository.Repository) Service {
	return service{repository}
}

type service struct {
	repository repository.Repository
}

func (s service) CreateBook(title string) (*models.Book, error) {
	book := models.Book{ID: bson.NewObjectId(), Title: title, CreatedAt: time.Now().UTC()}
	if err := s.repository.CreateBook(book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (s service) GetBook(id bson.ObjectId) (*models.Book, error) {
	return s.repository.GetBook(id)
}
