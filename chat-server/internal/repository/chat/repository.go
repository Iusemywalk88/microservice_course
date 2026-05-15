package chat

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/client/db"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/converter"
	sq "github.com/Masterminds/squirrel"
)

const (
	chatsTable      = "chats"
	idColumn        = "id"
	createdAtColumn = "created_at"

	chatMembersTable = "chat_members"
	chatIdColumn     = "chat_id"
	userIdColumn     = "user_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, req *model.Chat) (int64, error) {
	crReq, err := converter.ToChatFromService(req)
	if err != nil {
		return 0, err
	}

	builderChat := sq.Insert(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(createdAtColumn).
		Values(sq.Expr("NOW()")).
		Suffix("RETURNING id")

	query, args, err := builderChat.ToSql()
	if err != nil {
		return 0, err
	}

	qCT := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var chatID int64
	err = r.db.DB().QueryRowContext(ctx, qCT, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	builderMembers := sq.Insert(chatMembersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIdColumn, userIdColumn)

	for _, userID := range crReq.UserIDs {
		builderMembers = builderMembers.Values(chatID, userID)
	}

	query, args, err = builderMembers.ToSql()
	if err != nil {
		return 0, err
	}

	qCM := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, qCM, args...)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderChat := sq.Delete(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderChat.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
