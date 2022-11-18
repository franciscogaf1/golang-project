package model

type Question struct {
	ID       *int64 `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
	UserID   *int64 `json:"userId,omitempty"`
}