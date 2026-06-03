package repo

import (
	"database/sql"

	"enderz.net/testcontainer-test/internal/data"
)

type Repo struct {
	models   *data.Models
	Blogpost BlogpostReaderWriter
}

func NewRepo(db *sql.DB, models *data.Models) *Repo {
	return &Repo{
		models:   models,
		Blogpost: NewBlogpostService(db, models),
	}
}
