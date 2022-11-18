package controller

import (
	"level7/go/src/app/service"
	"net/http"

	"github.com/gorilla/mux"
)

var usersService = service.NewUsersService()

func SetupUsersEndpoints(r *mux.Router) {
	r.HandleFunc("/users", usersService.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{userId}", usersService.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/users/{userId}/questions", usersService.GetUserQuestionsByUserId).Methods(http.MethodGet)
	r.HandleFunc("/users", usersService.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{userId}", usersService.UpdateUser).Methods(http.MethodPatch, http.MethodPut)
	r.HandleFunc("/users/{userId}", usersService.DeleteUser).Methods(http.MethodDelete)
}
