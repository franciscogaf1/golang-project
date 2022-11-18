package controller

import (
	"level7/go/src/app/service"
	"net/http"

	"github.com/gorilla/mux"
)

var questionsService = service.NewQuestionsService()

func SetupQuestionsEndpoints(r *mux.Router) {
	r.HandleFunc("/questions", questionsService.GetAllQuestions).Methods(http.MethodGet)
	r.HandleFunc("/questions/{questionId}", questionsService.GetQuestionById).Methods(http.MethodGet)
	r.HandleFunc("/questions", questionsService.AddQuestion).Methods(http.MethodPost)
	r.HandleFunc("/questions/{questionId}", questionsService.UpdateQuestion).Methods(http.MethodPatch, http.MethodPut)
	r.HandleFunc("/questions/{questionId}", questionsService.DeleteQuestion).Methods(http.MethodDelete)
}