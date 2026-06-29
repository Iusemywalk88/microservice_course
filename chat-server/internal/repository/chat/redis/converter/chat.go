package converter

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
	modelRepo "github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat/redis/model"
	"strconv"
	"strings"
	"time"
)

func intToStringConverter(chat *model.Chat) string {
	var convertedNum string
	var stringSlice []string
	for _, i := range chat.UserIDs {
		convertedNum = strconv.Itoa(int(i))
		stringSlice = append(stringSlice, convertedNum)
	}
	return strings.Join(stringSlice, ",")
}

func stringToIntConverter(chat *modelRepo.Chat) []int64 {
	var intSlice []int64

	for _, s := range strings.Split(chat.UserIDs, ",") {
		if num, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
			intSlice = append(intSlice, int64(num))
		}
	}

	return intSlice
}

func ToRepoFromChat(chat *model.Chat) *modelRepo.Chat {
	return &modelRepo.Chat{
		ID:        chat.ID,
		CreatedAt: chat.CreatedAt.Unix(),
		UserIDs:   intToStringConverter(chat),
	}
}

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:        chat.ID,
		CreatedAt: time.Unix(chat.CreatedAt, 0),
		UserIDs:   stringToIntConverter(chat),
	}
}
