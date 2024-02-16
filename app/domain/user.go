package domain

import (
	"context"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetById(ctx context.Context, id int64) (User, error)
}