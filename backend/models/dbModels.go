package models

import (
	"time"

	"github.com/google/uuid"
)

type UserPrivate struct {
	UserId         uuid.UUID `json:"user_id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	RoleId         uuid.UUID `json:"role_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserPublic struct {
	Email string `json:"email"`
}
