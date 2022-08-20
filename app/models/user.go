package models

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	Rating    float32
	Role      int
	CreatedAt time.Time
}

func (u *User) CreateUSer() (err error) {
	createUser := `insert into users (
		uuid.
		name,
		email,
		password,
		rating,
		role,
		created_at
	) values(?, ?, ?, ? ,?, ?, ?)`

	_, err = Db.Exec(createUser, CreateUUID(), u.Name, u.Email, Encrypt(u.Password), u.Rating, u.Role, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	getUser := `select * from users where id = ?`
	err = Db.QueryRow(getUser, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
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
