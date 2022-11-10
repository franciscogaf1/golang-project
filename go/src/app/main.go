package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"level7/go/src/app/util"
	"level7/go/src/app/service"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var dbHost string = "postgres"
var dbPort string = "5432"
var dbName string = "postgres"
var dbUser string = "postgres"
var dbPassword string = "postgres"

func main() {

	connStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	util.SetDbConnection(db)

	r := mux.NewRouter()
	r.HandleFunc("/users", service.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{userId}", service.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/users/{userId}/questions", service.GetUserQuestionsByUserId).Methods(http.MethodGet)
	r.HandleFunc("/users", service.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{userId}", service.UpdateUser).Methods(http.MethodPatch, http.MethodPut)
	r.HandleFunc("/users/{userId}", service.DeleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/questions", service.GetAllQuestions).Methods(http.MethodGet)
	r.HandleFunc("/questions/{questionId}", service.GetQuestionById).Methods(http.MethodGet)
	r.HandleFunc("/questions", service.AddQuestion).Methods(http.MethodPost)
	r.HandleFunc("/questions/{questionId}", service.UpdateQuestion).Methods(http.MethodPatch, http.MethodPut)
	r.HandleFunc("/questions/{questionId}", service.DeleteQuestion).Methods(http.MethodDelete)
	fmt.Println("api running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}