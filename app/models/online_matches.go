package models

import (
	"time"
)

type OnlineMatch struct {
	ID                   int    `json:"id"`
	MaxParticipateNumber int    `json:"max_participate_number"`
	Status               string `json:"status"` // opening, processing, finishied
	// RemainingWaitTime  float64   `json:"remaining_wait_time"`
	QuestionNumber          int                     `json:"question_number"`
	WithBot                 bool                    `json:"with_bot"`
	RoomID                  *string                 `json:"room_id"`
	RoomPassword            *string                 `json:"room_password"`
	StartedAt               *time.Time              `json:"started_at"`
	FinishedAt              *time.Time              `json:"finished_at"`
	CreatedAt               *time.Time              `json:"created_at"`
	UpdatedAt               *time.Time              `json:"updated_at"`
	OnlineMatchJoinedUsers  []OnlineMatchJoinedUser `json:"online_match_joined_users"`
	OnlineMatchAskedQuizzes []OnlineMatchAskedQuiz  `json:"online_match_asked_quizzes"`
}

func GetOnlineMatchByID(id int) (online_match OnlineMatch, err error) {
	cmd := `select id,max_participate_number,started_at,finished_at,status,question_number,with_bot,room_id,room_password,created_at,updated_at from online_matches where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&online_match.ID,
		&online_match.MaxParticipateNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.QuestionNumber,
		&online_match.WithBot,
		&online_match.RoomID,
		&online_match.RoomPassword,
		&online_match.CreatedAt,
		&online_match.UpdatedAt,
	)

	return online_match, err
}

func GetOnlineMatchByRoomID(roomID string) (online_match OnlineMatch, err error) {
	cmd := `select id,max_participate_number,started_at,finished_at,status,question_number,with_bot,room_id,room_password,created_at,updated_at from online_matches where room_id = ? and status = ?`
	err = Db.QueryRow(cmd, roomID, "opening").Scan(
		&online_match.ID,
		&online_match.MaxParticipateNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.QuestionNumber,
		&online_match.WithBot,
		&online_match.RoomID,
		&online_match.RoomPassword,
		&online_match.CreatedAt,
		&online_match.UpdatedAt,
	)

	return online_match, err
}

func GetJoinableOnlineMatch() (online_match OnlineMatch, err error) {
	cmd := `select online_matches.id, online_matches.max_participate_number, online_matches.started_at, online_matches.finished_at, online_matches.status, online_matches.question_number, online_matches.with_bot, online_matches.room_id, online_matches.room_password, online_matches.created_at, online_matches.updated_at from online_matches left join online_match_joined_users on online_match_joined_users.online_match_id = online_matches.id where online_matches.status = ? group by online_matches.id having count(online_matches.id) < ?`

	err = Db.QueryRow(cmd, "opening", 4).Scan(
		&online_match.ID,
		&online_match.MaxParticipateNumber,
		&online_match.StartedAt,
		&online_match.FinishedAt,
		&online_match.Status,
		&online_match.QuestionNumber,
		&online_match.WithBot,
		&online_match.RoomID,
		&online_match.RoomPassword,
		&online_match.CreatedAt,
		&online_match.UpdatedAt,
	)

	return online_match, err
}

func (o *OnlineMatch) CreateOnlineMatch() (err error) {
	cmd := `insert into online_matches (
		room_id,
		room_password,
		with_bot,
		question_number,
		max_participate_number,
		status,
		created_at,
		updated_at
	) values(?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := Db.Exec(cmd, o.RoomID, o.RoomPassword, o.WithBot, o.QuestionNumber, o.MaxParticipateNumber, o.Status, time.Now(), time.Now())
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	o.ID = int(id)

	return err
}

func (m *OnlineMatch) UpdateOnlineMatch() (err error) {
	cmd, err := Db.Prepare(`update online_matches set status = ?,
	                                                  started_at = ?,
																										updated_at = ?,
																										finished_at = ?
																										where id = ?`)
	if err != nil {
		return err
	}
	_, err = cmd.Exec(m.Status, time.Now(), time.Now(), m.FinishedAt, m.ID)

	return err
}
