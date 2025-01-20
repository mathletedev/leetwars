package db

import (
	"context"

	"github.com/google/uuid"
)

func (d *Database) CreateUser(email string) {
	d.Query(
		context.Background(),
		"INSERT INTO users (id, email) VALUES ($1, $2);",
		uuid.NewString(),
		email,
	)
}
