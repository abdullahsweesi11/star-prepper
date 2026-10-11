package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId         uuid.UUID `db:"user_id"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
	RoleId         uuid.UUID `db:"role_id"`
	CreatedAt      time.Time `db:"created_at"`
}

type UserPublic struct {
	Email string `json:"email"`
}

type UserSecret struct {
	UserId         uuid.UUID `db:"user_id"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
}
