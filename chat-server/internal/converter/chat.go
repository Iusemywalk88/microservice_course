package converter

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
	"strconv"
)

func ToServiceFromDesc(req *desc.CreateRequest) (*model.Chat, error) {
	users := make([]int64, 0)
	for _, userID := range req.GetUsernames() {
		userID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return &model.Chat{}, err
		}
		users = append(users, userID)
	}

	return &model.Chat{
		UserIDs: users,
	}, nil
}
