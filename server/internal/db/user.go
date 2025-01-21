package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mathletedev/leetwars/internal/models"
)

func (d *Database) CreateUser(email string) (string, error) {
	id := uuid.NewString()

	rows, err := d.Query(
		context.Background(),
		"INSERT INTO users (id, email) VALUES ($1, $2);",
		id,
		email,
	)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	return id, nil
}

func (d *Database) ReadUser(id string) (*models.User, error) {
	rows, err := d.Query(
		context.Background(),
		"SELECT email, username FROM users WHERE id=$1;",
		id,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, err
	}

	var user models.User
	err = rows.Scan(&user.Email, &user.Username)
	if err != nil {
		return nil, errors.New("no username")
	}

	return &user, nil
}
