package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"level7/go/src/app/model"
	"level7/go/src/app/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type QuestionsRepoMock struct{}

func NewQuestionRepoMock() repository.QuestionsRepository {
	return &QuestionsRepoMock{}
}

func (*QuestionsRepoMock) FindAllQuestions() (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "question", "answer", "user_id"}).AddRow("1", "question1", "answer1", "1")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	rows, _ := db.Query("select")
	return rows, nil
}

func (*QuestionsRepoMock) FindQuestionById(questionId string) (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "question", "answer", "user_id"}).AddRow("1", "question1", "answer1", "1")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	rows, _ := db.Query("select")
	return rows, nil
}

func (*QuestionsRepoMock) InsertQuestion(question *model.Question, newId int64) *sql.Row {
	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id"}).AddRow("1")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	row := db.QueryRow("select")
	return row
}

func (*QuestionsRepoMock) UpdateQuestion(question *model.Question, questionId string) error {
	return nil
}

func (*QuestionsRepoMock) DeleteQuestion(questionId string) error {
	return nil
}

func (*QuestionsRepoMock) FindQuestionsByUserId(userId string) (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "question", "answer", "user_id"}).AddRow("1", "question1", "answer1", "1")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	rows, _ := db.Query("select")
	return rows, nil
}

var questionsService = QuestionsService{ QuestionsRepo: NewQuestionRepoMock() }

func TestGetAllQuestions(t *testing.T) {

	request, err := http.NewRequest("GET", "/questions", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(questionsService.GetAllQuestions)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	questions := []model.Question{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &questions)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, []model.Question{}, questions)

}

func TestGetQuestionById(t *testing.T) {

	request, err := http.NewRequest("GET", "/questions/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(questionsService.GetQuestionById)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	question := model.Question{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &question)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, model.Question{}, question)

}

func TestAddQuestion(t *testing.T) {

	payload := []byte(`{"question":"question","answer":"answer","user_id":"1"}`)

	request, err := http.NewRequest("POST", "/questions", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(questionsService.AddQuestion)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	question := model.Question{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &question)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, model.Question{}, question)

}

func TestUpdateQuestion(t *testing.T) {

	payload := []byte(`{"question":"different question","answer":"different answer"}`)

	request, err := http.NewRequest("PUT", "/questions/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(questionsService.UpdateQuestion)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	question := model.Question{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &question)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, model.Question{}, question)

}

func TestDeleteQuestion(t *testing.T) {

	request, err := http.NewRequest("DELETE", "/questions/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(questionsService.DeleteQuestion)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)
	assert.Empty(t, recorder.Body)

}