package data

import (
	"database/sql"
	"time"
)

type Models struct {
	Users    UserModel
	Blogpost BlogpostModel
	Groups   GroupModel
}

func NewModels(db *sql.DB, timeout *time.Duration) Models {
	return Models{
		Users:    UserModel{DB: db, Timeout: timeout},
		Blogpost: BlogpostModel{DB: db, Timeout: timeout},
		Groups:   GroupModel{DB: db, Timeout: timeout},
	}
}
