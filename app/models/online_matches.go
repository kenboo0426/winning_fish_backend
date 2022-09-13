package models

import (
	"log"
	"time"
)

type OnlineMatch struct {
	ID                 int       `json:"id"`
	PersonNumber       int       `json:"person_number"`
	ParticipantsNumber int       `json:"participants_number"`
	StartedAt          time.Time `json:"started_at"`
	FinishedAt         time.Time `json:"finished_at"`
	Status             string    `json:"status"` // opening, processing, finishied
	// RemainingWaitTime  float32   `json:"remaining_wait_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetOnlineMatch(id int) (online_match OnlineMatch, err error) {
	cmd := `select * from online_matches where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&online_match.ID,
		&online_match.PersonNumber,
		&online_match.ParticipantsNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.CreatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return online_match, err
}

func (o *OnlineMatch) CreateOnlineMatch() (err error) {
	cmd := `insert into online_match (
		person_number,
		participants_number,
		status,
		created_at,
		updated_at
	) values(?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, 0, 4, "opening", time.Now(), time.Now())

	if err != nil {
		log.Fatalln()
	}

	return err
}

func (m *OnlineMatch) UpdateOnlineMatch() (err error) {
	cmd, err := Db.Prepare(`update online_matches set status = ?,
	                                                  started_at = ?,
																										updated_at = ? 
																										where id = ?`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = cmd.Exec(m.Status, time.Now(), time.Now(), m.ID)

	return err
}
