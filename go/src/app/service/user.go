package service

import (
	"encoding/json"
	"io"
	"level7/go/src/app/util"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   *int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	rows, err := util.DBConn.Query("SELECT * FROM users")
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	users := make([]User, 0)

	for rows.Next() {
		var id *int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		users = append(users, User{ID: id, Name: name})
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(request)["userId"]

	rows, err := util.DBConn.Query("SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	user := User{}

	if rows.Next() {
		var id *int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		user.ID = id
		user.Name = name
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)

}

func GetUserQuestionsByUserId(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(request)["userId"]

	rows, err := util.DBConn.Query("SELECT q.id, q.question, q.answer, q.user_id FROM questions q JOIN users u ON q.user_id = u.id WHERE u.id = $1", userId)
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

func AddUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	user := User{}
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &user)
		if err != nil {
			log.Fatal(err)
			http.Error(writer, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		}

		var newId int64 = 0

		err = util.DBConn.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&newId)
		if err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		user.ID = &newId

		writer.WriteHeader(http.StatusAccepted)
		json.NewEncoder(writer).Encode(user)

	} else {
		http.Error(writer, "JSON Syntax Error", http.StatusNotAcceptable)
	}

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(request)["userId"]

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	user := User{}

	rows, err := util.DBConn.Query("SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	if rows.Next() {
		var id *int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		user.ID = id
		user.Name = name
	}

	if json.Valid(requestBody) {
		err = json.Unmarshal(requestBody, &user)
		if err != nil {
			log.Fatal(err)
			http.Error(writer, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		}

		_, err = util.DBConn.Exec("UPDATE users SET name = $1 WHERE id = $2", user.Name, userId)
		if err != nil {
			util.ThrowInternalServerError(&writer, err)
		}

		writer.WriteHeader(http.StatusAccepted)
		json.NewEncoder(writer).Encode(user)
	} else {
		http.Error(writer, "JSON Syntax Error", http.StatusNotAcceptable)
	}

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(request)["userId"]

	rows, err := util.DBConn.Query("SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		util.ThrowInternalServerError(&writer, err)
	}

	if rows.Next() {
		_, err = util.DBConn.Exec("DELETE FROM users WHERE id = $1", userId)
		if err != nil {
			util.ThrowInternalServerError(&writer, err)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}
}