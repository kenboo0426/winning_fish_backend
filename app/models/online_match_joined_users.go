package models

import (
	"log"
	"time"
)

type OnlineMatchJoinedUser struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	OnlineMatchID     int       `json:"online_match_id"`
	Rank              *int      `json:"rank"`
	RemainedTime      *float64  `json:"remained_time"`
	MissAnsweredCount *int      `json:"miss_answered_count"`
	Score             *int      `json:"score"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	User              User      `json:"user"`
}

func (o *OnlineMatchJoinedUser) CalculateRemainedTimeByOnlineMatchID() (err error) {
	isRemainedTime, _ := OnlineMatchJoinedUserHasRemainedTime(o.OnlineMatchID, o.UserID)

	if isRemainedTime {
		return err
	}

	total_remained_time, err := CalculateTotalRemainedTime(o.OnlineMatchID, o.UserID)
	if err != nil {
		log.Fatalln(err)
	}
	o.RemainedTime = &total_remained_time
	err = o.UpdateOnlineMatchJoinedUser()
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (o *OnlineMatchJoinedUser) UpdateOnlineMatchJoinedUser() (err error) {
	cmd, err := Db.Prepare(`update online_match_joined_users set remained_time = ?,
																															 miss_answered_count = ?,
																															 score = ?,
																															 updated_at = ?
																										           where id = ?`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = cmd.Exec(o.RemainedTime, o.MissAnsweredCount, o.Score, time.Now(), o.ID)

	return err
}

func (j *OnlineMatchJoinedUser) CreateOnlineMatchJoinedUser(userID string, onlineMatchID int) (err error) {
	_, err = GetJoinedUsersByOnlineMatchAndUserID(onlineMatchID, userID)
	if err == nil {
		return
	}

	cmd := `insert into online_match_joined_users (
			user_id,
			online_match_id,
			created_at,
			updated_at
		) values (?, ?, ?, ?)`
	result, err := Db.Exec(cmd, userID, onlineMatchID, time.Now(), time.Now())

	id, _ := result.LastInsertId()
	j.ID = int(id)

	return err
}

func OnlineMatchJoinedUserHasRemainedTime(online_match_id int, user_id int) (isRemainedTime bool, err error) {
	cmd := `select id, user_id, online_match_id, created_at from online_match_joined_users where online_match_id = ? and user_id = ? and remained_time is null`
	online_match_joined_user := OnlineMatchJoinedUser{}
	err = Db.QueryRow(cmd, online_match_id, user_id).Scan(
		&online_match_joined_user.ID,
		&online_match_joined_user.UserID,
		&online_match_joined_user.OnlineMatchID,
		&online_match_joined_user.CreatedAt,
	)

	if err != nil {
		return true, err
	} else {
		return false, err
	}

}

func GetJoinedUsersByOnlineMatchAndUserID(online_match_id int, user_id string) (online_match_joined_user OnlineMatchJoinedUser, err error) {
	cmd := `select id, user_id, online_match_id, created_at from online_match_joined_users where online_match_id = ? and user_id = ?`
	online_match_joined_user = OnlineMatchJoinedUser{}
	err = Db.QueryRow(cmd, online_match_id, user_id).Scan(
		&online_match_joined_user.ID,
		&online_match_joined_user.UserID,
		&online_match_joined_user.OnlineMatchID,
		&online_match_joined_user.CreatedAt,
	)

	return online_match_joined_user, err
}

func (o *OnlineMatch) GetJoinedUsersByOnlineMatch() (online_match_joined_users []OnlineMatchJoinedUser, err error) {
	cmd_to_fetch_general := `select t1.id, t1.user_id, t1.online_match_id, t1.rank, t1.remained_time, t1.miss_answered_count, t2.name, t2.icon from online_match_joined_users as t1 inner join users as t2 on t2.id = t1.user_id where t1.online_match_id = ?`
	rows, err := Db.Query(cmd_to_fetch_general, o.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var joinedUser OnlineMatchJoinedUser
		err = rows.Scan(&joinedUser.ID, &joinedUser.UserID, &joinedUser.OnlineMatchID, &joinedUser.Rank, &joinedUser.RemainedTime, &joinedUser.MissAnsweredCount, &joinedUser.User.Name, &joinedUser.User.Icon)
		if err != nil {
			log.Fatalln(err)
		}
		online_match_joined_users = append(online_match_joined_users, joinedUser)
	}
	rows.Close()

	return online_match_joined_users, err
}
