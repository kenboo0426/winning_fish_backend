package models

import "time"

type OnlineMatch struct {
	ID                 int       `json:"id"`
	PersonNumber       int       `json:"person_number"`
	ParticipantsNumber int       `json:"participants_number"`
	StartedAt          time.Time `json:"started_at"`
	FinishedAt         time.Time `json:"finished_at"`
	Status             string    `json:"status"` // opening, processing, finishied
	// RemainingWaitTime  float32   `json:"remaining_wait_time"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
