package model

import "time"

type Chat struct {
	ID        int64     // chats
	CreatedAt time.Time //chats
	UserIDs   []int64   //chat_members
}
