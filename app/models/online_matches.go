package models

import "time"

type OnlineMatch struct {
	ID                 int       `json:"id"`
	PersonNumber       int       `json:"person_number"`
	ParticipantsNumber int       `json:"participants_number"`
	StartedAt          time.Time `json:"started_at"`
	FinishedAt         time.Time `json:"finished_at"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
