package model

import (
	"database/sql"
	"time"
)

type Role int32

const (
	RoleUser  Role = 0
	RoleAdmin Role = 1
)

type User struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	UserRole  Role         `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
