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

type QuestionsServiceInterface interface {
	GetAllQuestions(writer http.ResponseWriter, request *http.Request)
	GetQuestionById(writer http.ResponseWriter, request *http.Request)
	AddQuestion(writer http.ResponseWriter, request *http.Request)
	UpdateQuestion(writer http.ResponseWriter, request *http.Request)
	DeleteQuestion(writer http.ResponseWriter, request *http.Request)
}

type QuestionsService struct {
	QuestionsRepo repository.QuestionsRepository
}

func NewQuestionsService() QuestionsServiceInterface {
	return &QuestionsService{repository.NewQuestionsRepo()}
}

func (service *QuestionsService) GetAllQuestions(writer http.ResponseWriter, request *http.Request) {
	rows, err := service.QuestionsRepo.FindAllQuestions()
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

func (service *QuestionsService) GetQuestionById(writer http.ResponseWriter, request *http.Request) {
	questionId := mux.Vars(request)["questionId"]

	rows, err := service.QuestionsRepo.FindQuestionById(questionId)
	if err != nil {
		http_util.ThrowInternalServerError(writer, err)
	}

	question := model.Question{}

	defer rows.Close()
	if rows.Next() {
		var id *int64
		var questionField string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &questionField, &answer, &userId); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		question.ID = id
		question.Question = questionField
		question.Answer = answer
		question.UserID = userId
	}

	http_util.SuccessResponse(writer, &question, http.StatusOK)
}

func (service *QuestionsService) AddQuestion(writer http.ResponseWriter, request *http.Request) {
	question := model.Question{}
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &question)
		if err != nil {
			http_util.ThrowBadRequestError(writer, err)
		}

		var newId int64 = 0

		row := service.QuestionsRepo.InsertQuestion(&question, newId)
		if err = row.Scan(&newId); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		question.ID = &newId

		http_util.SuccessResponse(writer, &question, http.StatusOK)
	} else {
		http_util.ThrowBadRequestError(writer, fmt.Errorf("syntax error"))
	}

}

func (service *QuestionsService) UpdateQuestion(writer http.ResponseWriter, request *http.Request) {
	questionId := mux.Vars(request)["questionId"]

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		http_util.ThrowBadRequestError(writer, err)
	}

	question := model.Question{}

	rows, err := service.QuestionsRepo.FindQuestionById(questionId)
	if err != nil {
		http_util.ThrowInternalServerError(writer, err)
	}

	defer rows.Close()
	if rows.Next() {
		var id *int64
		var questionField string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &questionField, &answer, &userId); err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		question.ID = id
		question.Question = questionField
		question.Answer = answer
		question.UserID = userId
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &question)
		if err != nil {
			http_util.ThrowBadRequestError(writer, err)
		}

		err = service.QuestionsRepo.UpdateQuestion(&question, questionId)
		if err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}

		http_util.SuccessResponse(writer, &question, http.StatusOK)
	} else {
		http_util.ThrowBadRequestError(writer, fmt.Errorf("syntax error"))
	}

}

func (service *QuestionsService) DeleteQuestion(writer http.ResponseWriter, request *http.Request) {
	questionId := mux.Vars(request)["questionId"]

	rows, err := service.QuestionsRepo.FindQuestionById(questionId)
	if err != nil {
		http_util.ThrowInternalServerError(writer, err)
	}

	defer rows.Close()
	if rows.Next() {
		err = service.QuestionsRepo.DeleteQuestion(questionId)
		if err != nil {
			http_util.ThrowInternalServerError(writer, err)
		}
		http_util.SuccessResponseNoBody(writer, http.StatusOK)
	}
}
