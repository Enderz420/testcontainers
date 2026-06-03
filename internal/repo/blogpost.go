package repo

import (
	"context"
	"database/sql"

	"enderz.net/testcontainer-test/internal/data"
	"github.com/google/uuid"
)

type BlogpostReader interface {
	Read(ctx context.Context, id uuid.UUID) (*data.Blogpost, error)
	List(ctx context.Context) ([]*data.Blogpost, error)
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

func newBlogpostService(db *sql.DB, models *data.Models) BlogpostReaderWriter {
	return &BlogpostService{
		db:     db,
		models: models,
	}
}

func (s *BlogpostService) Read(ctx context.Context, id uuid.UUID) (*data.Blogpost, error) {
	return nil, nil
}

func (s *BlogpostService) List(ctx context.Context) ([]*data.Blogpost, error) {
	return nil, nil
}

func (s *BlogpostService) Create(ctx context.Context, input data.BlogpostInput) (*data.Blogpost, error) {
	return nil, nil
}

func (s *BlogpostService) Update(ctx context.Context, id uuid.UUID, input data.BlogpostPatch) (*data.Blogpost, error) {
	return nil, nil
}

func (s *BlogpostService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
