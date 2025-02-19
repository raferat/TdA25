package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

func scanUserFromRow[T interface{ Scan(...any) error }](r T) (*User, error) {
	elem := &User{}
	err := r.Scan(&elem.Uuid,
		&elem.CreatedAt.Time,
		&elem.Username,
		&elem.Email,
		&elem.Password,
		&elem.Elo,
		&elem.Wins,
		&elem.Draws,
		&elem.Losses,
		&elem.Banned,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error scanning user values: %#v\n", err)
		}
		return nil, err
	}

	return elem, nil
}

func CreateUser(base *UserBase) *User {
	user := &User{
		UserBase:  *base,
		Uuid:      uuid.NewString(),
		CreatedAt: MyTime{Time: time.Now()},
		Wins:      0,
		Draws:     0,
		Losses:    0,
		Banned:    false,
	}

	_, err := createUser.Exec(user.Uuid, user.CreatedAt.Time, user.Username, user.Email, user.Password, user.Elo, user.Wins, user.Draws, user.Losses, user.Banned)
	if err != nil {
		log.Printf("Error creating user: %#v\n", err)
		return nil
	}

	return user
}

func ListUsers() (result []*User) {
	rows, err := listUsers.Query()
	if err != nil {
		log.Printf("Error listing users: %#v\n", err)
		return nil
	}

	for rows.Next() {
		usr, err := scanUserFromRow(rows)
		if err != nil {
			log.Printf("Error retrieving use info: %#v\n", err)
			return nil
		}
		result = append(result, usr)
	}

	return
}

func GetUserByUUID(uuid string) *User {
	row := findUser.QueryRow(uuid)
	user, err := scanUserFromRow(row)
	if err != nil {
		return nil
	}
	return user
}

func UpdateUser(uuid string, base *UserBase) *User {
	user := &User{
		UserBase: *base,
		Uuid:     uuid,
	}
	err := editUser.QueryRow(base.Username, base.Email, base.Password, base.Elo, uuid).Scan(&user.CreatedAt.Time, &user.Wins, &user.Draws, &user.Losses, &user.Banned)
	if err != nil {
		log.Printf("Error updating user: %#v\n", err)
		return nil
	}

	return user
}

// true if successful
func DeleteUser(uuid string) bool {
	state, err := deleteUser.Exec(uuid)
	if err != nil {
		log.Printf("Error deleting user: %#v\n", err)
		return false
	}

	num, err := state.RowsAffected()
	if err != nil {
		log.Println("Toto nesmi nastat. Pokud toto nastalo, znamena to konec civilizace. Buh nam vsem pomahej.")
		panic(err)
	}

	return num == 1
}
