package models

import "time"

type Answer struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	QuizID           int       `json:"quiz_id"`
	Correct          bool      `json:"correct"`
	AnsweredOptionID int       `json:"answered_option_id"`
	CreatedAt        time.Time `json:"created_at"`
}
