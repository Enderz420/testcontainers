package repo

import (
	"context"
	"database/sql"
	"log/slog"

	"enderz.net/testcontainer-test/internal/data"
	"enderz.net/testcontainer-test/internal/database"
	"enderz.net/testcontainer-test/internal/logging"
	"github.com/google/uuid"
)

type UserReader interface {
	Read(ctx context.Context, id uuid.UUID) (*data.User, error)
	List(ctx context.Context) ([]*data.User, *database.Metadata, error)
}

type UserWriter interface {
	Create(ctx context.Context, input data.UserInput) (*data.User, error)
	Update(ctx context.Context, input data.UserPatch) (*data.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserReaderWriter interface {
	UserReader
	UserWriter
}

type UserService struct {
	db     *sql.DB
	models *data.Models
}

func NewUserService(db *sql.DB, models *data.Models) UserReaderWriter {
	return &UserService{
		db:     db,
		models: models,
	}
}

func (r *UserService) Read(ctx context.Context, id uuid.UUID) (*data.User, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"readUser", slog.String("id", id.String())))

	logger.LogAttrs(ctx, slog.LevelInfo, "reading user")
	user, err := r.models.Users.SelectOne(ctx, id)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "user read successfully", slog.String("id", user.ID.String()))

	return user, nil
}

func (r *UserService) List(ctx context.Context) ([]*data.User, *database.Metadata, error) {

	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"listUsers"))

	logger.LogAttrs(ctx, slog.LevelInfo, "listing users")
	users, metadata, err := r.models.Users.SelectAll(ctx)
	if err != nil {
		return nil, nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "users listed successfully", slog.Int64("length", metadata.Length))

	return users, metadata, nil
}

func (r *UserService) Create(ctx context.Context, input data.UserInput) (*data.User, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"createUser", slog.Any("input", input)))

	logger.LogAttrs(ctx, slog.LevelInfo, "creating user")

	user, err := r.models.Users.Insert(ctx, input)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "user created successfully", slog.Any("id", user.ID))

	return user, nil
}

func (r *UserService) Update(ctx context.Context, input data.UserPatch) (*data.User, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"updateUser", slog.Any("input", input)))

	logger.LogAttrs(ctx, slog.LevelInfo, "updating user")

	user, err := r.models.Users.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "user updated successfully", slog.Any("id", user.ID))

	return user, nil
}

func (r *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"deleteUser", slog.String("id", id.String())))

	logger.LogAttrs(ctx, slog.LevelInfo, "deleting user")

	err := r.models.Users.Delete(ctx, id)
	if err != nil {
		return err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "user deleted successfully", slog.Any("id", id))

	return nil
}
