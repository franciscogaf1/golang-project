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

type UsersRepoMock struct{}

func NewUsersRepoMock() repository.UsersRepository {
	return &UsersRepoMock{}
}

func (*UsersRepoMock) FindAllUsers() (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "name"}).AddRow("1", "francisco")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	rows, _ := db.Query("select")
	return rows, nil
}

func (*UsersRepoMock) FindUserById(userId string) (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "name"}).AddRow("1", "francisco")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	rows, _ := db.Query("select")
	return rows, nil
}

func (*UsersRepoMock) InsertUser(user *model.User, newId int64) *sql.Row {
	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id"}).AddRow("1")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	row := db.QueryRow("select")
	return row
}

func (*UsersRepoMock) UpdateUser(user *model.User, userId string) error {
	return nil
}

func (*UsersRepoMock) DeleteUser(userId string) error {
	return nil
}

var usersService = UsersService{ QuestionsRepo: NewQuestionRepoMock(), UsersRepo: NewUsersRepoMock() }

func TestGetAllUsers(t *testing.T) {

	request, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(usersService.GetAllUsers)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	users := []model.User{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, []model.User{}, users)

}

func TestGetUserById(t *testing.T) {

	request, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(questionsService.GetQuestionById)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	user := model.User{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, model.User{}, user)

}

func TestAddUser(t *testing.T) {

	payload := []byte(`{"name":"francisco"}`)

	request, err := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(usersService.AddUser)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	user := model.User{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, model.User{}, user)

}

func TestUpdateUser(t *testing.T) {

	payload := []byte(`{"name":"goncalves"}`)

	request, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(usersService.UpdateUser)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)

	user := model.User{}

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, model.User{}, user)

}

func TestDeleteUser(t *testing.T) {

	request, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(usersService.DeleteUser)
	handler.ServeHTTP(recorder, request)

	assert.EqualValues(t, recorder.Code, http.StatusOK)
	assert.NotNil(t, recorder.Body)
	assert.Empty(t, recorder.Body)

}

func TestGetUserQuestionsByUserId(t *testing.T) {

	request, err := http.NewRequest("GET", "/users/1/questions", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(usersService.GetUserQuestionsByUserId)
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
