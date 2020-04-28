package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chidiwilliams/go-web-server-tips/server/handlers"
	"github.com/chidiwilliams/go-web-server-tips/server/handlers/books"
	"github.com/chidiwilliams/go-web-server-tips/services/book"
	"github.com/chidiwilliams/go-web-server-tips/services/book/repository"
)

var (
	bookHandler books.BookHandler
)

// Server configures and returns a new http.Server
func Server() *http.Server {
	r := mux.NewRouter()

	r.Handle("/book", handlers.Handler(bookHandler.CreateBook)).Methods(http.MethodPost)
	r.Handle("/book/{bookID}", handlers.Handler(bookHandler.GetBook)).Methods(http.MethodGet)

	srv := &http.Server{Handler: r, Addr: ":8080"}
	return srv
}

func init() {
	inMemoryDB, err := connectToInMemoryDB()
	fatalIfErr(err)

	mongoDB, err := connectToMongo()
	fatalIfErr(err)

	err = ensureMongoIndexes(mongoDB)
	fatalIfErr(err)

	// Switch repository to Mongo with:
	// _ = repository.NewInMemoryRepository(inMemoryDB)
	// bookRepository := repository.NewMongoRepository(mongoDB)
	bookRepository := repository.NewInMemoryRepository(inMemoryDB)
	_ = repository.NewMongoRepository(mongoDB)

	bookService := book.NewService(bookRepository)
	bookHandler = books.NewBookHandler(bookService)
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
