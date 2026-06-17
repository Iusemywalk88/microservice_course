package converter

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
	modelRepo "github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat/pg/model"
)

func ToChatFromService(req *model.Chat) (*modelRepo.Chat, error) {
	return &modelRepo.Chat{
		ID:        req.ID,
		CreatedAt: req.CreatedAt,
		UserIDs:   req.UserIDs,
	}, nil
}
