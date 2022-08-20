package models

import "time"

type Option struct {
	ID        int
	Name      string
	QuizID    int
	CreatedAt time.Time
}
