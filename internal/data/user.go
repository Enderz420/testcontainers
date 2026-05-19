package data

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"
	"time"

	"enderz.net/testcontainer-test/internal/apperrors"
	"enderz.net/testcontainer-test/internal/logging"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

type User struct {
	ID          mssql.UniqueIdentifier `json:"id"`
	Username    string                 `json:"username"`
	Email       string                 `json:"email"`
	Password    string                 `json:"-"`
	CreatedAt   time.Time              `json:"created_at"`
	LastUpdated time.Time              `json:"last_updated"`
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
INSERT INTO User (
	id, username, email, password, created_at, last_updated)
VALUES (
	DEFAULT, $1, $2, $3, GETUTCDATE(), GETUTCDATE()
)
RETURNING id, username, password, created_at, last_updated;
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
		us.Username,
		us.Email,
		us.Password,
		time.Now(),
		time.Now(),
	).Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.LastUpdated,
	)
	if err != nil {
		if strings.Contains(err.Error(), "unique_violation") || strings.Contains(err.Error(), "23505") {
			return nil, ErrDuplicateUsername
		}
		return nil, err
	}

	return &result, nil
}

func (m UserModel) GetAll(ctx context.Context) ([]*User, error) {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
SELECT id, username, email, password, created_at, last_updated
FROM User
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
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.LastUpdated,
		); err != nil {
			logger.ErrorContext(ctx, "error scanning row", "error", err)
			return nil, err
		}
		results = append(results, &user)
	}

	if err := rows.Err(); err != nil {
		logger.ErrorContext(ctx, "error iterating rows", "error", err)
		return nil, err
	}

	return results, nil
}

func (m UserModel) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
SELECT CAST(id AS CHAR(36)), username, email, password, created_at, last_updated
FROM User
WHERE id = $1;
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
	err := m.DB.QueryRowContext(ctx, stmt, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.LastUpdated,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			logger.InfoContext(ctx, "no rows found")
			return nil, apperrors.ErrRecordNotFound
		default:
			logger.ErrorContext(ctx, "error scanning row", "error", err)
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) Delete(ctx context.Context, userID uuid.UUID) error {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
DELETE FROM User
WHERE id = $1;
`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger = logger.With(
		slog.Group(
			"query",
			slog.String("statement", stmt),
			slog.String("id", userID.String()),
		),
	)

	logger.InfoContext(ctx, "performing query")

	_, err := m.DB.ExecContext(ctx, stmt, userID)
	if err != nil {
		logger.ErrorContext(ctx, "error deleting user", "error", err)
		return err
	}

	logger.Info("user deleted successfully")
	return nil
}
