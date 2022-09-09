package models

import "time"

type OnlineMatchJoinedUser struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	Rank              int       `json:"rank"`
	RemainedTime      float32   `json:"remained_time"`
	MissAnsweredCount int       `json:"miss_answered_count"`
	Score             int       `json:"score"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
