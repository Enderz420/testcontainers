package data

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"
	"time"

	"enderz.net/testcontainer-test/internal/database"
	"enderz.net/testcontainer-test/internal/logging"
	"enderz.net/testcontainer-test/internal/models"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

type User struct {
	ID        mssql.UniqueIdentifier `json:"id"`
	Username  string                 `json:"username"`
	Email     string                 `json:"email"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type UserModel struct {
	DB      *sql.DB
	Timeout *time.Duration
}

var (
	ErrDuplicateUsername = errors.New("duplicate username")
	ErrDuplicateEmail    = errors.New("duplicate email")
)

func (m UserModel) Insert(ctx context.Context, us *User) (*User, error) {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
INSERT INTO [User] (
	id,
	username,
	email,
	created_at,
	updated_at)
OUTPUT
	INSERTED.id,
	INSERTED.username,
	INSERTED.email,
	INSERTED.created_at,
	INSERTED.updated_at
VALUES (
	NEWID(), @Username, @Email, GETDATE(), GETDATE()
)
`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	newUUID := uuid.New()
	us.ID = mssql.UniqueIdentifier(newUUID)

	var result User

	logger = logger.With(
		slog.Group(
			"query",
			slog.String("statement", stmt),
			"user", us,
		),
	)

	logger.InfoContext(ctx, "performing query")

	err := m.DB.QueryRowContext(
		ctx,
		stmt,
		sql.Named("Username", us.Username),
		sql.Named("Email", us.Email),
	).Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "unique_violation") || strings.Contains(err.Error(), "23505") {
			return nil, ErrDuplicateUsername
		}
		return nil, err
	}

	return &result, nil
}

func (m UserModel) SelectAll(ctx context.Context) ([]*User, *database.Metadata, error) {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
SELECT id, username, email, created_at, updated_at
FROM [User]
ORDER BY id DESC;
`

	logger = logger.With(
		slog.Group(
			"query",
			slog.String("statement", stmt),
		),
	)

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	var results []*User

	logger.InfoContext(ctx, "performing query")

	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			logger.ErrorContext(ctx, "error scanning row", "error", err)
			return nil, nil, err
		}
		results = append(results, &user)
	}

	if err := rows.Err(); err != nil {
		logger.ErrorContext(ctx, "error iterating rows", "error", err)
		return nil, nil, err
	}

	metadata := database.NewMetadata(results)
	if metadata.Length > 0 {
		metadata.LastSeen = uuid.UUID(results[metadata.Length-1].ID)
	}
	logger.Info("query successful", slog.Any("metadata", metadata))

	return results, &metadata, nil
}

func (m UserModel) SelectOne(ctx context.Context, id mssql.UniqueIdentifier) (*User, error) {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
SELECT id, username, email, created_at, updated_at
FROM [User]
WHERE id = @ID;
`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	var user User

	logger = logger.With(
		slog.Group(
			"query",
			slog.String("statement", stmt),
			slog.String("id", id.String()),
		),
	)

	logger.InfoContext(ctx, "performing query")
	err := m.DB.QueryRowContext(ctx, stmt, sql.Named("ID", id)).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			logger.InfoContext(ctx, "no rows found")
			return nil, models.ErrRecordNotFound
		default:
			logger.ErrorContext(ctx, "error scanning row", "error", err)
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) Delete(ctx context.Context, id mssql.UniqueIdentifier) error {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
DELETE FROM [User]
WHERE id = @ID;
`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger = logger.With(
		slog.Group(
			"query",
			slog.String("statement", stmt),
			slog.String("id", id.String()),
		),
	)

	logger.InfoContext(ctx, "performing query")

	_, err := m.DB.ExecContext(ctx, stmt, sql.Named("ID", id))
	if err != nil {
		logger.ErrorContext(ctx, "error deleting user", "error", err)
		return err
	}

	logger.Info("user deleted successfully")
	return nil
}
