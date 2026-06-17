package model

type Chat struct {
	ID        int64   `redis:"id"`         // chats
	CreatedAt int64   `redis:"created_at"` //chats
	UserIDs   []int64 //chat_members
}
