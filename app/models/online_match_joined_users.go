package models

import (
	"log"
	"time"
)

type OnlineMatchJoinedUser struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	OnlineMatchID     int       `json:"online_match_id"`
	Rank              int       `json:"rank"`
	RemainedTime      float32   `json:"remained_time"`
	MissAnsweredCount int       `json:"miss_answered_count"`
	Score             int       `json:"score"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (j *OnlineMatchJoinedUser) CreateOnlineMatchJoinedUser(userID string, onlineMatchID string) (err error) {
	cmd := `insert into online_match_joined_users (
		user_id,
		online_match_id,
		created_at,
		updated_at,
	) values (?,?,?,?)`

	result, err := Db.Exec(cmd, userID, onlineMatchID, time.Now(), time.Now())
	id, _ := result.LastInsertId()
	j.ID = int(id)

	return err
}

func (o *OnlineMatch) GetJoinedUsersByOnlineMatch() (online_match_joined_users []OnlineMatchJoinedUser, err error) {
	cmd := `select id, user_id, online_match_id from online_match_joined_users where online_match_id = ?`
	rows, err := Db.Query(cmd, o.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var joinedUser OnlineMatchJoinedUser
		err = rows.Scan(&joinedUser.ID, &joinedUser.UserID, &joinedUser.OnlineMatchID)
		if err != nil {
			log.Fatalln(err)
		}
		online_match_joined_users = append(online_match_joined_users, joinedUser)
	}
	rows.Close()

	return online_match_joined_users, err
}
