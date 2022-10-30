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
	Rating    *float64  `json:"rating"`
	Role      int       `json:"role"` // 0: general 1: admin 2: guests
	Icon      *string   `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CreateUser() (err error) {
	createUser := `insert into users (
		uuid,
		name,
		email,
		role,
		icon,
		created_at
	) values(?, ?, ?, ?, ?, ?)`

	result, err := Db.Exec(createUser, u.UUID, u.Name, u.Email, u.Role, u.Icon, time.Now())
	id, _ := result.LastInsertId()
	u.ID = int(id)

	return err
}

func GetUserByIDOrUUID(id int, uuid string) (user User, err error) {
	user = User{}
	getUser := `select id, uuid, name, email, rating, role, created_at, icon from users where id = ? or uuid = ?`
	err = Db.QueryRow(getUser, id, uuid).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Rating,
		&user.Role, // 0: general, 1: admin, 2: guest
		&user.CreatedAt,
		&user.Icon,
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
