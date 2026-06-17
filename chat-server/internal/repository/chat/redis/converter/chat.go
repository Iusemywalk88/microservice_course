package converter

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat/pg/model"
	modelRepo "github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat/redis/model"
)

func ToRepoFromChat(chat *model.Chat) *modelRepo.Chat {
	return &modelRepo.Chat{
		ID:        chat.ID,
		CreatedAt: chat.CreatedAt.Unix(),
		UserIDs:   chat.UserIDs,
	}
}
