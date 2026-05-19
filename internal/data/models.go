package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound       = errors.New("record not found")
	ErrConstraintViolation  = errors.New("constraint violation")
	ErrUniqueIndexViolation = errors.New("unique violation")
	ErrUniqueKeyViolation   = errors.New("unique key violation")
)

type Models struct {
	Users UserModel
}

func NewModels(db *sql.DB, timeout *time.Duration) Models {
	return Models{
		Users: UserModel{DB: db, Timeout: timeout},
	}
}
