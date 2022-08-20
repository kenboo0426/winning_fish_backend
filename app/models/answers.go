package models

import "time"

type Answer struct {
	ID               int
	UserID           int
	QuizID           int
	Correct          bool
	AnsweredOptionID int
	CreatedAt        time.Time
}
