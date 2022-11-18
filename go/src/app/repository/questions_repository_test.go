package repository

import (
	"database/sql"
	"level7/go/src/app/model"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewQuestionRepoMock(db *sql.DB) QuestionsRepository {
	return &QuestionsRepo{db}
}

func TestFindAllQuestions(t *testing.T) {

	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "question", "answer", "user_id"}).AddRow("1", "question1", "answer1", "1")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM questions`)).WillReturnRows(mockRows)

	testQuestionsRepo := NewQuestionRepoMock(db)

	rows, err := testQuestionsRepo.FindAllQuestions()
	if err != nil {
		log.Fatal(err)
	}

	questions := []model.Question{}

	for rows.Next() {
		var id *int64
		var question string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &question, &answer, &userId); err != nil {
			t.Fatal(err)
		}

		questions = append(questions, model.Question{ID: id, Question: question, Answer: answer, UserID: userId})
	}

	expectedQuestions := []model.Question{}
	var questionId int64 = 1
	var userId int64 = 1
	question := model.Question{ID: &questionId, Question: "question1", Answer: "answer1", UserID: &userId}
	expectedQuestions = append(expectedQuestions, question)

	assert.NotNil(t, questions)
	assert.EqualValues(t, expectedQuestions, questions)
}

func TestFindQuestionById(t *testing.T) {

	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "question", "answer", "user_id"}).AddRow("1", "question1", "answer1", "1")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM questions WHERE id = $1`)).WillReturnRows(mockRows)

	testQuestionsRepo := NewQuestionRepoMock(db)

	rows, err := testQuestionsRepo.FindQuestionById("1")
	if err != nil {
		log.Fatal(err)
	}

	questions := []model.Question{}

	for rows.Next() {
		var id *int64
		var question string
		var answer string
		var userId *int64

		if err := rows.Scan(&id, &question, &answer, &userId); err != nil {
			t.Fatal(err)
		}

		questions = append(questions, model.Question{ID: id, Question: question, Answer: answer, UserID: userId})
	}

	expectedQuestions := []model.Question{}
	var questionId int64 = 1
	var userId int64 = 1
	question := model.Question{ID: &questionId, Question: "question1", Answer: "answer1", UserID: &userId}
	expectedQuestions = append(expectedQuestions, question)

	assert.NotNil(t, questions)
	assert.EqualValues(t, expectedQuestions, questions)

}

func TestInsertQuestion(t *testing.T) {

	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO questions (question, answer, user_id) VALUES ($1, $2, $3) RETURNING id`)).WillReturnRows(mockRows)

	testQuestionsRepo := NewQuestionRepoMock(db)

	var userId int64 = 1
	question := model.Question{ID: nil, Question: "question", Answer: "answer", UserID: &userId}
	var actualId int64 = 0

	row := testQuestionsRepo.InsertQuestion(&question, actualId)
	if err := row.Scan(actualId); err != nil {
		log.Fatal(err)
	}

	var expectedId int64 = 1

	assert.EqualValues(t, expectedId, actualId)

}

func TestUpdateQuestion(t *testing.T) {

	db, mock, _ := sqlmock.New()
	mockRows := sqlmock.NewRows([]string{"id", "question", "answer", "user_id"}).AddRow("1", "question1", "answer1", "1")
	mock.ExpectQuery(regexp.QuoteMeta(`"UPDATE questions SET question = $1, answer = $2 WHERE id = $3"`)).WillReturnRows(mockRows)

	testQuestionsRepo := NewQuestionRepoMock(db)

	var questionId int64 = 1
	var userId int64 = 1
	question := model.Question{ID: &questionId, Question: "aaaa", Answer: "bbbb", UserID: &userId}

	err := testQuestionsRepo.UpdateQuestion(&question, "1")
	if err != nil {
		log.Fatal(err)
	}

	//assert.NotNil(t, questions)
	//assert.EqualValues(t, expectedQuestions, questions)

}

func TestDeleteQuestion(t *testing.T) {

	
}

func TestFindQuestionsByUserId(t *testing.T) {


}
