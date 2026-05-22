package data

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"enderz.net/testcontainer-test/internal/database"
	"enderz.net/testcontainer-test/internal/logging"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

type BlogpostModel struct {
	DB      *sql.DB
	Timeout *time.Duration
}

type Blogpost struct {
	ID        mssql.UniqueIdentifier `json:"id"`
	Title     string                 `json:"title"`
	Content   string                 `json:"content"`
	CreatedBy string				 `json:"created_by"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type BlogpostInput struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedBy string 	`json:"created_by"`
}

func (m BlogpostModel) Insert(ctx context.Context, input BlogpostInput) (*Blogpost, error) {
	logger := logging.LoggerFromContext(ctx)

	const stmt string = `
	INSERT INTO Blogpost (
		id,
		title,
		content,
		created_by,
		created_at,
		updated_at)
	OUTPUT
		INSERTED.id,
		INSERTED.title,
		INSERTED.content,
		INSERTED.created_by,
		INSERTED.created_at,
		INSERTED.updated_at
	VALUES (NEWID(), @Title, @Content, @Createdby, GETDATE(), GETDATE())
	`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger.Info("performing query")

	var blogpost Blogpost
	row := m.DB.QueryRowContext(
		ctx,
		stmt,
		sql.Named("Title", input.Title),
		sql.Named("Content", input.Content),
		sql.Named("Createdby", input.CreatedBy),
	)

	err := row.Scan(
		&blogpost.ID,
		&blogpost.Title,
		&blogpost.Content,
		&blogpost.CreatedBy,
		&blogpost.CreatedAt,
		&blogpost.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	logger.Info("query successful")

	return &blogpost, nil
}

func (m BlogpostModel) SelectOne(ctx context.Context, id mssql.UniqueIdentifier) (*Blogpost, error) {
	logger := logging.LoggerFromContext(ctx)

	const stmt string = `
	SELECT id, title, content, created_by, created_at, updated_at
	FROM Blogpost
	WHERE id = @p1
	`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger.Info("performing query")

	var blogpost Blogpost
	row := m.DB.QueryRowContext(
		ctx,
		stmt,
		id,
	)

	err := row.Scan(
		&blogpost.ID,
		&blogpost.Title,
		&blogpost.Content,
		&blogpost.CreatedBy,
		&blogpost.CreatedAt,
		&blogpost.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	logger.Info("query successfull")

	return &blogpost, nil
}
func (m BlogpostModel) SelectAll(ctx context.Context) ([]*Blogpost, *database.Metadata, error) {
	logger := logging.LoggerFromContext(ctx)

	const stmt string = `
	SELECT id, title, content, created_by, created_at, updated_at
	FROM Blogpost
	ORDER BY "id"
	`

	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	logger.Info("performing query")

	var results []*Blogpost
	rows, err := m.DB.QueryContext(
		ctx,
		stmt,
	)
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		var blogpost Blogpost
		err = rows.Scan(
			&blogpost.ID,
			&blogpost.Title,
			&blogpost.Content,
			&blogpost.CreatedBy,
			&blogpost.CreatedAt,
			&blogpost.UpdatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		results = append(results, &blogpost)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	metadata := database.NewMetadata(results)
	if metadata.Length > 0 {
		metadata.LastSeen = uuid.UUID(results[metadata.Length-1].ID)
	}
	logger.Info("query successful", slog.Any("metadata", metadata))

	return results, &metadata, nil
}

func (m BlogpostModel) Delete(ctx context.Context, id mssql.UniqueIdentifier) error {
	logger := logging.LoggerFromContext(ctx)

	stmt := `
	DELETE FROM Blogpost
	WHERE id = @p1;
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

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		logger.ErrorContext(ctx, "error deleting blogpost", "error", err)
		return err
	}

	logger.Info("blogpost deleted successfully")
	return nil
}