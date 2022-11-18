package service

import (
	"encoding/json"
	"fmt"
	"io"
	"level7/go/src/app/http_util"
	"level7/go/src/app/model"
	"level7/go/src/app/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UsersServiceInterface interface {
	GetAllUsers(writer http.ResponseWriter, request *http.Request)
	GetUserById(writer http.ResponseWriter, request *http.Request)
	AddUser(writer http.ResponseWriter, request *http.Request)
	GetUserQuestionsByUserId(writer http.ResponseWriter, request *http.Request)
	UpdateUser(writer http.ResponseWriter, request *http.Request)
	DeleteUser(writer http.ResponseWriter, request *http.Request)
}

type UsersService struct {
	QuestionsRepo repository.QuestionsRepository
	UsersRepo     repository.UsersRepository
}

func NewUsersService() UsersServiceInterface {
	return &UsersService{
		repository.NewQuestionsRepo(),
		repository.NewUsersRepo(),
	}
}

func (service *UsersService) GetAllUsers(writer http.ResponseWriter, request *http.Request) {

	rows, err := service.UsersRepo.FindAllUsers()
	if err != nil {
		http_util.ThrowBadRequestError(writer, err)
	}

	users := make([]model.User, 0)

	defer rows.Close()
	for rows.Next() {
		var id *int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		users = append(users, model.User{ID: id, Name: name})
	}

	http_util.SuccessResponse(writer, &users, http.StatusOK)
}

func (service *UsersService) GetUserById(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["userId"]

	rows, err := service.UsersRepo.FindUserById(userId)
	if err != nil {
		http_util.ThrowInternalServerError(writer, err)
	}

	user := model.User{}

	defer rows.Close()
	if rows.Next() {
		var id *int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		user.ID = id
		user.Name = name
	}

	http_util.SuccessResponse(writer, &user, http.StatusOK)

}

func (service *UsersService) GetUserQuestionsByUserId(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["userId"]

	rows, err := service.QuestionsRepo.FindQuestionsByUserId(userId)
	if err != nil {
		http_util.ThrowInternalServerError(writer, err)
	}

	questions := make([]model.Question, 0)

	defer rows.Close()
	for rows.Next() {
		var id *int64
		var question string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &question, &answer, &userId); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		questions = append(questions, model.Question{ID: id, Question: question, Answer: answer, UserID: userId})
	}

	http_util.SuccessResponse(writer, &questions, http.StatusOK)
}

func (service *UsersService) AddUser(writer http.ResponseWriter, request *http.Request) {

	user := model.User{}
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &user)
		if err != nil {
			http_util.ThrowBadRequestError(writer, err)
		}

		var newId int64 = 0

		row := service.UsersRepo.InsertUser(&user, newId)
		if err = row.Scan(&newId); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		user.ID = &newId

		http_util.SuccessResponse(writer, &user, http.StatusOK)
	} else {
		http_util.ThrowBadRequestError(writer, fmt.Errorf("syntax error"))
	}

}

func (service *UsersService) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["userId"]

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	user := model.User{}

	rows, err := service.UsersRepo.FindUserById(userId)
	if err != nil {
		http_util.ThrowInternalServerError(writer, err)
	}

	defer rows.Close()
	if rows.Next() {
		var id *int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		user.ID = id
		user.Name = name
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &user)
		if err != nil {
			http_util.ThrowBadRequestError(writer, err)
		}

		err = service.UsersRepo.UpdateUser(&user, userId)
		if err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		http_util.SuccessResponse(writer, &user, http.StatusOK)
	} else {
		http_util.ThrowBadRequestError(writer, fmt.Errorf("syntax error"))
	}

}

func (service *UsersService) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["userId"]

	rows, err := service.UsersRepo.FindUserById(userId)
	if err != nil {
		http_util.ThrowInternalServerError(writer, err)
	}

	defer rows.Close()
	if rows.Next() {
		err = service.UsersRepo.DeleteUser(userId)
		if err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}
		http_util.SuccessResponseNoBody(writer, http.StatusOK)
	}
}
