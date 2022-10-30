package models

import (
	"log"
	"time"
)

type OnlineMatch struct {
	ID                   int        `json:"id"`
	MaxParticipantNumber int        `json:"max_participant_number"`
	StartedAt            *time.Time `json:"started_at"`
	FinishedAt           *time.Time `json:"finished_at"`
	Status               string     `json:"status"` // opening, processing, finishied
	// RemainingWaitTime  float64   `json:"remaining_wait_time"`
	QuestionNumber          int                     `json:"question_number"`
	WithBot                 bool                    `json:"with_bot"`
	RoomID                  *string                 `json:"room_id"`
	CreatedAt               *time.Time              `json:"created_at"`
	UpdatedAt               *time.Time              `json:"updated_at"`
	OnlineMatchJoinedUsers  []OnlineMatchJoinedUser `json:"online_match_joined_users"`
	OnlineMatchAskedQuizzes []OnlineMatchAskedQuiz  `json:"online_match_asked_quizzes"`
}

func GetOnlineMatch(id int) (online_match OnlineMatch, err error) {
	cmd := `select id,max_participant_number,started_at,finished_at,status,question_number,with_bot,room_id,created_at,updated_at from online_matches where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&online_match.ID,
		&online_match.MaxParticipantNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.QuestionNumber,
		&online_match.WithBot,
		&online_match.RoomID,
		&online_match.CreatedAt,
		&online_match.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return online_match, err
}

func GetJoinableOnlineMatch() (online_match OnlineMatch, err error) {
	cmd := `select t1.id, t1.participants_number, t1.started_at, t1.finished_at, t1.status, t1.question_number, t1.with_bot, t1.room_id, t1.created_at, t1.updated_at from t1 left join online_match_joined_users on online_match_joined_users.online_match_id = t1.id where t1.status = ? group by t1.id having count(t1.id) < ?`

	err = Db.QueryRow(cmd, "opening", 4).Scan(
		&online_match.ID,
		&online_match.MaxParticipantNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.QuestionNumber,
		&online_match.WithBot,
		&online_match.RoomID,
		&online_match.CreatedAt,
		&online_match.UpdatedAt,
	)

	return online_match, err
}

func (o *OnlineMatch) CreateOnlineMatch() (err error) {
	cmd := `insert into online_matches (
		max_participant_number,
		status,
		created_at,
		updated_at
	) values(?, ?, ?, ?, ?)`

	result, err := Db.Exec(cmd, o.MaxParticipantNumber, o.Status, time.Now(), time.Now())
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
																										updated_at = ?,
																										finished_at = ?
																										where id = ?`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = cmd.Exec(m.Status, time.Now(), time.Now(), m.FinishedAt, m.ID)

	return err
}
