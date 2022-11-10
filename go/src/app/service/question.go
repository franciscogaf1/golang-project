package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"level7/go/src/app/util"

	"github.com/gorilla/mux"
)

type Question struct {
	ID       *int64 `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
	UserID   *int64 `json:"userId,omitempty"`
}

// Question Related Functions
func GetAllQuestions(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	rows, err := util.DBConn.Query("SELECT * FROM questions")
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	questions := make([]Question, 0)

	for rows.Next() {
		var id *int64
		var question string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &question, &answer, &userId); err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		questions = append(questions, Question{ID: id, Question: question, Answer: answer, UserID: userId})
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(questions)
}

func GetQuestionById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	questionId := mux.Vars(request)["questionId"]

	rows, err := util.DBConn.Query("SELECT * FROM questions WHERE id = $1", questionId)
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	question := Question{}

	if rows.Next() {
		var id *int64
		var questionField string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &questionField, &answer, &userId); err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		question.ID = id
		question.Question = questionField
		question.Answer = answer
		question.UserID = userId
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(question)

}

func AddQuestion(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	question := Question{}
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &question)
		if err != nil {
			log.Fatal(err)
			http.Error(writer, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		}

		var newId int64 = 0

		err = util.DBConn.QueryRow("INSERT INTO questions (question, answer, user_id) VALUES ($1, $2, $3) RETURNING id", question.Question, question.Answer, *question.UserID).Scan(&newId)
		if err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		question.ID = &newId

		writer.WriteHeader(http.StatusAccepted)
		json.NewEncoder(writer).Encode(question)

	} else {
		http.Error(writer, "JSON Syntax Error", http.StatusNotAcceptable)
	}

}

func UpdateQuestion(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	questionId := mux.Vars(request)["questionId"]

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	question := Question{}

	rows, err := util.DBConn.Query("SELECT * FROM questions WHERE id = $1", questionId)
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	if rows.Next() {
		var id *int64
		var questionField string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &questionField, &answer, &userId); err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		question.ID = id
		question.Question = questionField
		question.Answer = answer
		question.UserID = userId
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &question)
		if err != nil {
			log.Fatal(err)
			http.Error(writer, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		}

		_, err = util.DBConn.Exec("UPDATE questions SET question = $1, answer = $2 WHERE id = $3", question.Question, question.Answer, questionId)
		if err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		writer.WriteHeader(http.StatusAccepted)
		json.NewEncoder(writer).Encode(question)
	} else {
		http.Error(writer, "JSON Syntax Error", http.StatusNotAcceptable)
	}
	
}

func DeleteQuestion(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	questionId := mux.Vars(request)["questionId"]

	rows, err := util.DBConn.Query("SELECT * FROM questions WHERE id = $1", questionId)
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	if rows.Next() {
		_, err = util.DBConn.Exec("DELETE FROM questions WHERE id = $1", questionId)
		if err != nil {
			util.ThrowInternalServerError(&writer, err)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}
}
