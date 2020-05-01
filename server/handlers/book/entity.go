package book

import "github.com/chidiwilliams/go-web-server-tips/models"

type getBookResponse struct {
	Book *models.Book `json:"book"`
}

type createBookResponse struct {
	Book *models.Book `json:"book"`
}
