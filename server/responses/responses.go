package responses

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Body       *responseBody
	StatusCode int
}

type responseBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ToJSON writes the response to the given http.ResponseWriter
// with an application/json Content-Type header.
func (r response) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	return json.NewEncoder(w).Encode(r.Body)
}

// OK returns a new successful response.
func OK(message string, data interface{}) *response {
	return newResponse(true, message, data, http.StatusOK)
}

// Fail returns a new failed response.
func Fail(message string, statusCode int) *response {
	return newResponse(false, message, nil, statusCode)
}

func newResponse(success bool, message string, data interface{}, statusCode int) *response {
	return &response{
		Body: &responseBody{
			Success: success,
			Message: message,
			Data:    data,
		},
		StatusCode: statusCode,
	}
}
