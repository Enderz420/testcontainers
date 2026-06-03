package data

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"enderz.net/testcontainer-test/internal/logging"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

type Group struct {
	ID           mssql.UniqueIdentifier `json:"id"`
	Name         string                 `json:"name"`
	CreatedAt    time.Time              `json:"created_at"`
	LastModified time.Time              `json:"last_modified"`
}

type GroupInsert struct {
	Name string `json:"name"`
}

type GroupModel struct {
	DB      *sql.DB
	Timeout *time.Duration
}

func (m *GroupModel) SelectOne(
	ctx context.Context,
	id string,
) (*Group, error) {

	const stmt = `
	SELECT id, name, created_at, last_modified
	FROM core.groups
	WHERE id = @ID
	`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger := logging.LoggerFromContext(ctx)

	logger.Info("performing query")
	var group Group
	row := m.DB.QueryRowContext(ctx, stmt, sql.Named("ID", id))

	if err := row.Scan(&group.ID, &group.Name, &group.CreatedAt, &group.LastModified); err != nil {
		return nil, err
	}

	return &group, nil
}

func (m *GroupModel) Insert(ctx context.Context, input GroupInsert) (*Group, error) {
	const stmt = `
	INSERT INTO core.groups (name)
	OUTPUT INSERTED.id, INSERTED.name, INSERTED.created_at, INSERTED.last_modified
	VALUES (@Name)
	`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger := logging.LoggerFromContext(ctx).With(
		slog.Group("query", slog.String("statement", stmt), "group", input),
	)

	var group Group
	logger.LogAttrs(ctx, slog.LevelInfo, "performing query")

	row := m.DB.QueryRowContext(ctx, stmt, sql.Named("Name", input.Name))
	if err := row.Scan(&group.ID, &group.Name, &group.CreatedAt, &group.LastModified); err != nil {
		return nil, err
	}

	return &group, nil
}

func (m *GroupModel) Delete(ctx context.Context, id uuid.UUID) error {
	const stmt = `
	DELETE FROM core.groups
	WHERE id = @ID
	`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger := logging.LoggerFromContext(ctx).With(
		slog.Group("query", slog.String("statement", stmt), "group", id),
	)

	logger.LogAttrs(ctx, slog.LevelInfo, "performing query")

	row := m.DB.QueryRowContext(ctx, stmt, sql.Named("ID", id))
	if err := row.Scan(); err != nil {
		return err
	}

	return nil
}
