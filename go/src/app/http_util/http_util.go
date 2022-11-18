package http_util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorMessage struct {
	Code    *int    `json:"code"`
	Message string `json:"message"`
}

// Success
func SuccessResponse[T any](writer http.ResponseWriter, t *T, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(t)
}

func SuccessResponseNoBody(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
}

// Exceptions
func ThrowInternalServerError(writer http.ResponseWriter, err error) {
	throwError(writer, err, http.StatusInternalServerError)
}

func ThrowBadRequestError(writer http.ResponseWriter, err error) {
	throwError(writer, err, http.StatusBadRequest)
}

func throwError(writer http.ResponseWriter, err error, statusCode int) {
	log.Fatal(err)
	http.Error(writer, http.StatusText(statusCode), statusCode)
	message := ErrorMessage{&statusCode, http.StatusText(statusCode)}
	json.NewEncoder(writer).Encode(message)
}
