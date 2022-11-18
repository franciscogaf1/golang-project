package repository

import (
	"database/sql"
	"level7/go/src/app/model"
)

type QuestionsRepository interface {
	FindAllQuestions() (*sql.Rows, error)
	FindQuestionById(questionId string) (*sql.Rows, error)
	FindQuestionsByUserId(userId string) (*sql.Rows, error)
	InsertQuestion(question *model.Question, newId int64) *sql.Row
	UpdateQuestion(question *model.Question, questionId string) error
	DeleteQuestion(questionId string) error
}

type QuestionsRepo struct{
	db *sql.DB
}

func NewQuestionsRepo() QuestionsRepository {
	return &QuestionsRepo{ GetDBConn() }
}

func (questionsRepo *QuestionsRepo) FindAllQuestions() (*sql.Rows, error) {
	rows, err := questionsRepo.db.Query("SELECT * FROM questions")
	return rows, err
}

func (questionsRepo *QuestionsRepo) FindQuestionById(questionId string) (*sql.Rows, error) {
	rows, err := questionsRepo.db.Query("SELECT * FROM questions WHERE id = $1", questionId)
	return rows, err
}

func (questionsRepo *QuestionsRepo) InsertQuestion(question *model.Question, newId int64) (*sql.Row) {
	row := questionsRepo.db.QueryRow("INSERT INTO questions (question, answer, user_id) VALUES ($1, $2, $3) RETURNING id", question.Question, question.Answer, *question.UserID)
	return row
}

func (questionsRepo *QuestionsRepo) UpdateQuestion(question *model.Question, questionId string) error {
	_, err := questionsRepo.db.Exec("UPDATE questions SET question = $1, answer = $2 WHERE id = $3", question.Question, question.Answer, questionId)
	return err
}

func (questionsRepo *QuestionsRepo) DeleteQuestion(questionId string) error {
	_, err := questionsRepo.db.Exec("DELETE FROM questions WHERE id = $1", questionId)
	return err
}

func (questionsRepo *QuestionsRepo) FindQuestionsByUserId(userId string) (*sql.Rows, error) {
	rows, err := questionsRepo.db.Query("SELECT q.id, q.question, q.answer, q.user_id FROM questions q JOIN users u ON q.user_id = u.id WHERE u.id = $1", userId)
	return rows, err
}
