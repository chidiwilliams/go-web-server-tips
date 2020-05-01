package book

import (
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"github.com/chidiwilliams/go-web-server-tips/errors"
	"github.com/chidiwilliams/go-web-server-tips/server/decoder"
	"github.com/chidiwilliams/go-web-server-tips/server/responses"
	"github.com/chidiwilliams/go-web-server-tips/services/book"
)

type Handler interface {
	CreateBook(w http.ResponseWriter, r *http.Request) error
	GetBook(w http.ResponseWriter, r *http.Request) error
}

func NewBookHandler(bookService book.Service) Handler {
	return handler{bookService}
}

type handler struct {
	bookService book.Service
}

type createBookRequestBody struct {
	Title string `json:"title"`
}

func (u handler) CreateBook(w http.ResponseWriter, r *http.Request) error {
	requestBody := &createBookRequestBody{}
	if err := decoder.DecodeJSON(r.Body, requestBody); err != nil {
		return err
	}

	newBook, err := u.bookService.CreateBook(requestBody.Title)
	if err != nil {
		return err
	}

	return responses.OK("We've added your book!", createBookResponse{Book: newBook}).ToJSON(w)
}

func (u handler) GetBook(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["bookID"]
	if !bson.IsObjectIdHex(id) {
		return errors.Error("invalid vendor ID")
	}

	retrievedBook, err := u.bookService.GetBook(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return responses.OK("We found your book!", getBookResponse{retrievedBook}).ToJSON(w)
}
