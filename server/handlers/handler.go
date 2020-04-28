package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	errors2 "github.com/chidiwilliams/go-web-server-tips/errors"
	"github.com/chidiwilliams/go-web-server-tips/server/responses"
)

// Handler is an implementation of http.Handler with a handler
// function that returns an error
type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err == nil {
		return
	}

	if err = respondWithErr(err, w); err != nil {
		log.Println("error writing http response:", err)
	}
}

func respondWithErr(err error, w http.ResponseWriter) error {
	appError := new(errors2.AppError)
	if errors.As(err, &appError) { // client error
		return responses.Fail(
			toSentenceCase(err.Error()),
			errTypeStatusCode(appError.Type()),
		).ToJSON(w)
	}

	log.Println("server error:", err)
	return responses.Fail("Internal Server Error", http.StatusInternalServerError).ToJSON(w)
}

func errTypeStatusCode(errType errors2.Type) int {
	switch errType {
	case errors2.TypeBadRequest:
		return http.StatusBadRequest
	case errors2.TypeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusBadRequest
	}
}

func toSentenceCase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
