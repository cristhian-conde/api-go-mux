package user

import (
	"awesomeProject/types"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func (store *Store) GetUserById(id int) (*types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (store *Store) CreateUser(user types.User) error {
	//TODO implement me
	panic("implement me")
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := store.db.Query("SELECT * email FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(row *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, err
}
