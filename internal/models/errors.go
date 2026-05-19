package models

import "errors"

var (
	ErrRecordNotFound       = errors.New("record not found")
	ErrConstraintViolation  = errors.New("constraint violation")
	ErrUniqueIndexViolation = errors.New("unique violation")
	ErrUniqueKeyViolation   = errors.New("unique key violation")
)
