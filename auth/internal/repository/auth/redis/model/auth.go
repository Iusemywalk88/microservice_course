package model

type Role int32

const (
	RoleUser  Role = 0
	RoleAdmin Role = 1
)

type User struct {
	ID        int64  `redis:"id"`
	Name      string `redis:"name"`
	Email     string `redis:"email"`
	Password  string `redis:"password"`
	UserRole  Role   `redis:"role"`
	CreatedAt int64  `redis:"created_at"`
	UpdatedAt *int64 `redis:"updated_at"`
}
