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

type UserGroupJunction struct {
	UserID  mssql.UniqueIdentifier
	GroupID mssql.UniqueIdentifier
}

type UserGroupModel struct {
	DB      *sql.DB
	Timeout *time.Duration
}

func (m *UserGroupModel) Insert(ctx context.Context, userID, groupID mssql.UniqueIdentifier) (*UserGroupJunction, error) {
	logger := logging.LoggerFromContext(ctx)

	const stmt = `
		INSERT INTO core.user_group_junction
			(user_id, group_id)
		OUTPUT
			INSERTED.user_id, INSERTED.group_id
		VALUES (@UserID, @GroupID)
	`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger.Info("inserting user group junction", "user_id", userID, "group_id", groupID)

	var junction UserGroupJunction

	row := m.DB.QueryRowContext(ctx, stmt, sql.Named("UserID", userID), sql.Named("GroupID", groupID))

	err := row.Scan(&junction.UserID, &junction.GroupID)
	if err != nil {
		return nil, err
	}

	logger.Info("inserted user group junction", "user_id", junction.UserID, "group_id", junction.GroupID)

	return &junction, nil
}

func (m *UserGroupModel) DeleteUserFromGroup(ctx context.Context, userID, groupID uuid.UUID) error {
	logger := logging.LoggerFromContext(ctx)

	const stmt = `
		DELETE FROM core.user_group_junction
		WHERE user_id = @UserID AND group_id = @GroupID
	`

	logger.Info("stmt", slog.Any("stmt", stmt))
	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger.Info("deleting user from group", "user_id", userID, "group_id", groupID)
	_, err := m.DB.ExecContext(ctx, stmt, sql.Named("UserID", userID), sql.Named("GroupID", groupID))
	if err != nil {
		return err
	}
	logger.Info("query successful")

	return nil
}

func (m *UserGroupModel) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	logger := logging.LoggerFromContext(ctx)

	const stmt = `
		DELETE FROM core.user_group_junction
		WHERE user_id = @UserID
	`

	logger.Info("stmt", stmt, slog.Info)

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger.InfoContext(ctx, "Performing query")

	_, err := m.DB.ExecContext(ctx, stmt, sql.Named("UserID", userID))
	if err != nil {
		return err
	}

	logger.Info("user deleted")

	return nil
}
