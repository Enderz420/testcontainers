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

type GroupReader interface {
	List(ctx context.Context) ([]*data.Group, *database.Metadata, error)
	Read(ctx context.Context, id uuid.UUID) (*data.Group, error)
}

type GroupWriter interface {
	Insert(ctx context.Context, group data.GroupInsert) (*data.Group, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type GroupReaderWriter interface {
	GroupReader
	GroupWriter
}

type GroupService struct {
	db     *sql.DB
	models *data.Models
}

func NewGroupService(db *sql.DB, models *data.Models) *GroupService {
	return &GroupService{
		db:     db,
		models: models,
	}
}

func (r *GroupService) List(ctx context.Context) ([]*data.Group, *database.Metadata, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"readGroups"))

	logger.LogAttrs(ctx, slog.LevelInfo, "reading groups")

	groups, metadata, err := r.models.Group.SelectAll(ctx)
	if err != nil {
		return nil, nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "groups read successfully", slog.Any("groups", groups))

	return groups, metadata, nil
}

func (r *GroupService) Read(ctx context.Context, id uuid.UUID) (*data.Group, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"readGroup", slog.String("id", id.String())))

	logger.LogAttrs(ctx, slog.LevelInfo, "reading group")

	group, err := r.models.Group.SelectOne(ctx, id)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "group read successfully", slog.Any("group", group))

	return group, nil
}

func (r *GroupService) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"deleteGroup", slog.String("id", id.String())))
	logger.LogAttrs(ctx, slog.LevelInfo, "deleting group")

	if err := r.models.Group.Delete(ctx, id); err != nil {
		return err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "group deleted successfully")

	return nil
}

func (r *GroupService) Insert(ctx context.Context, input data.GroupInsert) (*data.Group, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group("insertGroup", slog.Any("group", input)))
	logger.LogAttrs(ctx, slog.LevelInfo, "inserting group")

	group, err := r.models.Group.Insert(ctx, input)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "group inserted successfully", slog.Any("group", group))

	return group, nil
}
