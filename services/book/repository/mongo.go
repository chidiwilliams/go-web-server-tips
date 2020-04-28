package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/chidiwilliams/go-web-server-tips/models"
)

func NewMongoRepository(db *mgo.Database) Repository {
	return mongoRepository{coll: db.C("books")}
}

type mongoRepository struct {
	coll *mgo.Collection
}

func (m mongoRepository) CreateBook(book models.Book) error {
	if err := m.coll.Insert(book); err != nil {
		return err
	}

	return nil
}

func (m mongoRepository) GetBook(id bson.ObjectId) (*models.Book, error) {
	book := new(models.Book)

	if err := m.coll.FindId(id).One(book); err != nil {
		if err == mgo.ErrNotFound {
			return nil, errBookNotFound
		}
		return nil, err
	}

	return book, nil
}
