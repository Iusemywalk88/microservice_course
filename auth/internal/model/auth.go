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
	ID        int64
	Name      string
	Email     string
	Password  string
	UserRole  Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UpdateUserInfo struct {
	ID       int64
	Name     *string
	Email    *string
	UserRole Role
}
