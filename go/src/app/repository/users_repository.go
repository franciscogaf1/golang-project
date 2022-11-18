package repository

import (
	"database/sql"
	"level7/go/src/app/model"
)

type UsersRepository interface {
	FindAllUsers() (*sql.Rows, error)
	FindUserById(userId string) (*sql.Rows, error)
	InsertUser(user *model.User, newId int64) *sql.Row
	UpdateUser(user *model.User, questionId string) error
	DeleteUser(userId string) error
}

type UsersRepo struct{
	db *sql.DB
}

func NewUsersRepo() UsersRepository {
	return &UsersRepo{ GetDBConn() }
}

func (usersRepo *UsersRepo) FindAllUsers() (*sql.Rows, error) {
	rows, err := usersRepo.db.Query("SELECT * FROM users")
	return rows, err
}

func (usersRepo *UsersRepo) FindUserById(userId string) (*sql.Rows, error) {
	rows, err := usersRepo.db.Query("SELECT * FROM users WHERE id = $1", userId)
	return rows, err
}

func (usersRepo *UsersRepo) InsertUser(user *model.User, newId int64) *sql.Row {
	row := usersRepo.db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name)
	return row
}

func (usersRepo *UsersRepo) UpdateUser(user *model.User, userId string) error {
	_, err := usersRepo.db.Exec("UPDATE users SET name = $1 WHERE id = $2", user.Name, userId)
	return err
}

func (usersRepo *UsersRepo) DeleteUser(userId string) error {
	_, err := usersRepo.db.Exec("DELETE FROM users WHERE id = $1", userId)
	return err
}
