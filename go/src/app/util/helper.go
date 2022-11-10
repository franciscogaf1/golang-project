package util

import (
	"database/sql"
	"log"
	"net/http"
)

var DBConn *sql.DB

// Expose DB
func SetDbConnection(db *sql.DB) {
	DBConn = db
}

// Helpers
func ThrowInternalServerError(writer *http.ResponseWriter, err error) {
	log.Fatal(err)
	http.Error(*writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
