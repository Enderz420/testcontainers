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

type BlogpostReader interface {
	Read(ctx context.Context, id uuid.UUID) (*data.Blogpost, error)
	List(ctx context.Context) ([]*data.Blogpost, *database.Metadata, error)
}

type BlogpostWriter interface {
	Create(ctx context.Context, input data.BlogpostInput) (*data.Blogpost, error)
	Update(ctx context.Context, id uuid.UUID, input data.BlogpostPatch) (*data.Blogpost, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type BlogpostReaderWriter interface {
	BlogpostReader
	BlogpostWriter
}

type BlogpostService struct {
	db     *sql.DB
	models *data.Models
}

func NewBlogpostService(db *sql.DB, models *data.Models) BlogpostReaderWriter {
	return &BlogpostService{
		db:     db,
		models: models,
	}
}

func (r *BlogpostService) Read(ctx context.Context, id uuid.UUID) (*data.Blogpost, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"readBlogpost", slog.String("id", id.String())))

	logger.LogAttrs(ctx, slog.LevelInfo, "reading blogpost")

	blogpost, err := r.models.Blogpost.SelectOne(ctx, id)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "blogpost read successfully", slog.Any("blogpost", blogpost))

	return blogpost, nil
}

func (r *BlogpostService) List(ctx context.Context) ([]*data.Blogpost, *database.Metadata, error) {
	return nil, nil, nil
}

func (r *BlogpostService) Create(ctx context.Context, input data.BlogpostInput) (*data.Blogpost, error) {
	return nil, nil
}

func (r *BlogpostService) Update(ctx context.Context, id uuid.UUID, input data.BlogpostPatch) (*data.Blogpost, error) {
	return nil, nil
}

func (r *BlogpostService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
