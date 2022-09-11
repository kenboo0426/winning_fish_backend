package models

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Rating    *float32  `json:"rating"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CreateUser() (err error) {
	createUser := `insert into users (
		uuid,
		name,
		email,
		role,
		created_at
	) values(?, ?, ? , ?, ?)`

	_, err = Db.Exec(createUser, u.UUID, u.Name, u.Email, 0, time.Now())


	return err
}

func GetUserByIDOrUUID(id int, uuid string) (user User, err error) {
	user = User{}
	getUser := `select * from users where id = ? or uuid = ?`
	err = Db.QueryRow(getUser, id, uuid).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Rating,
		&user.Role,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	updateUser := `update users set name = ?, email = ?, role = ?, rating = ? where id = ?`
	_, err = Db.Exec(updateUser, u.Name, u.Email, u.Rating, u.Role, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	deleteUser := `delete from users where id = ?`
	_, err = Db.Exec(deleteUser, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
