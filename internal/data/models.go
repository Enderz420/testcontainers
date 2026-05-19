package data

import (
	"database/sql"
	"time"
)

type Models struct {
	Users UserModel
}

func NewModels(db *sql.DB, timeout *time.Duration) Models {
	return Models{
		Users: UserModel{DB: db, Timeout: timeout},
	}
}
