package models

import (
	"log"
	"time"
)

type OnlineMatch struct {
	ID                 int        `json:"id"`
	PersonNumber       int        `json:"person_number"`
	ParticipantsNumber int        `json:"participants_number"`
	StartedAt          *time.Time `json:"started_at"`
	FinishedAt         *time.Time `json:"finished_at"`
	Status             string     `json:"status"` // opening, processing, finishied
	// RemainingWaitTime  float32   `json:"remaining_wait_time"`
	CreatedAt              *time.Time              `json:"created_at"`
	UpdatedAt              *time.Time              `json:"updated_at"`
	OnlineMatchJoinedUsers []OnlineMatchJoinedUser `json:"online_match_joined_users"`
}

func GetOnlineMatch(id int) (online_match OnlineMatch, err error) {
	cmd := `select id,person_number,participants_number,started_at,finished_at,status,created_at,updated_at from online_matches where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&online_match.ID,
		&online_match.PersonNumber,
		&online_match.ParticipantsNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.CreatedAt,
		&online_match.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return online_match, err
}

func GetJoinableOnlineMatch() (online_match OnlineMatch, err error) {
	cmd := `select online_matches.id, online_matches.person_number, online_matches.participants_number, online_matches.started_at, online_matches.finished_at, online_matches.status, online_matches.created_at, online_matches.updated_at from online_matches left join online_match_joined_users on online_match_joined_users.online_match_id = online_matches.id where online_matches.status = ? group by online_matches.id having count(online_matches.id) < ?`

	err = Db.QueryRow(cmd, "opening", 4).Scan(
		&online_match.ID,
		&online_match.PersonNumber,
		&online_match.ParticipantsNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.CreatedAt,
		&online_match.UpdatedAt,
	)

	return online_match, err
}

func (o *OnlineMatch) CreateOnlineMatch() (err error) {
	cmd := `insert into online_matches (
		person_number,
		participants_number,
		status,
		created_at,
		updated_at
	) values(?, ?, ?, ?, ?)`

	result, err := Db.Exec(cmd, o.PersonNumber, o.ParticipantsNumber, o.Status, time.Now(), time.Now())
	id, _ := result.LastInsertId()
	o.ID = int(id)

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
