package model

import "time"

type Chat struct {
	ID        int64
	CreatedAt time.Time
	UserIDs   []int64
}
