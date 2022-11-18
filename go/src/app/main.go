package main

import (
	"level7/go/src/app/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	controller.SetupUsersEndpoints(r)
	controller.SetupQuestionsEndpoints(r)

	log.Println("api running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
