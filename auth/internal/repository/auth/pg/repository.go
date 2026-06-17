package pg

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/db"
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository"
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/pg/converter"
	modelRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/pg/model"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, req *model.User) (int64, error) {
	crReq := converter.ToRepoFromUser(req)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn, passwordColumn).
		Values(crReq.Name, crReq.Email, crReq.UserRole, crReq.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "auth_repository.Get", QueryRaw: query}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Update(ctx context.Context, user *model.UpdateUserInfo) error {
	updateReq := converter.ToRepoFromServiceUpdate(user)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(updatedAtColumn, sq.Expr("NOW()")).
		Where(sq.Eq{idColumn: updateReq.ID})

	if updateReq.Name != nil {
		builder = builder.Set(nameColumn, *updateReq.Name)
	}
	if updateReq.Email != nil {
		builder = builder.Set(emailColumn, *updateReq.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
