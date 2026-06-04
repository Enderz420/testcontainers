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
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"listBlogpost"))

	logger.LogAttrs(ctx, slog.LevelInfo, "listing blogposts")

	blogposts, metadata, err := r.models.Blogpost.SelectAll(ctx)
	if err != nil {
		return nil, nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "blogposts listed successfully", slog.Int("count", len(blogposts)))

	return blogposts, metadata, nil
}

func (r *BlogpostService) Create(ctx context.Context, input data.BlogpostInput) (*data.Blogpost, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"createBlogpost", slog.Any("input", input)))

	logger.LogAttrs(ctx, slog.LevelInfo, "creating blogpost")
	blogpost, err := r.models.Blogpost.Insert(ctx, input)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "blogpost created successfully", slog.Any("blogpost", blogpost))

	return blogpost, nil
}

func (r *BlogpostService) Update(ctx context.Context, id uuid.UUID, input data.BlogpostPatch) (*data.Blogpost, error) {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"updateBlogpost", slog.Any("input", input)))

	logger.LogAttrs(ctx, slog.LevelInfo, "updating blogpost")
	blogpost, err := r.models.Blogpost.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "blogpost updated successfully", slog.Any("blogpost", blogpost))

	return blogpost, nil
}

func (r *BlogpostService) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, logger := logging.ContextLogger(ctx, slog.Group(
		"deleteBlogpost", slog.Any("id", id)))

	logger.LogAttrs(ctx, slog.LevelInfo, "deleting blogpost")
	err := r.models.Blogpost.Delete(ctx, id)
	if err != nil {
		return err
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "blogpost deleted successfully", slog.Any("id", id))

	return nil
}
