package models

import "time"

type OnlineMatchAskedQuiz struct {
	ID            int       `json:"id"`
	QuizID        int       `json:"quiz_id"`
	OnlineMatchID int       `json:"online_match_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
