package data

import (
	"database/sql"
	"time"
)

type Models struct {
	Users    UserModel
	Blogpost BlogpostModel
	Group    GroupModel
}

func NewModels(db *sql.DB, timeout *time.Duration) Models {
	return Models{
		Users:    UserModel{DB: db, Timeout: timeout},
		Blogpost: BlogpostModel{DB: db, Timeout: timeout},
		Group:    GroupModel{DB: db, Timeout: timeout},
	}
}
