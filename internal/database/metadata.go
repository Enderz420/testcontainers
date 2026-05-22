package database

import (
	"reflect"

	"github.com/google/uuid"
)

type Metadata struct {
	LastSeen uuid.UUID `json:"last_seen,omitzero"`
	Length   int64     `json:"length,omitzero"`
}

// CC: Copied from regnskapsportalen
// NewMetadata uses query results and filter to create a new [Metadata] instance.
//
// The following is required by the input:
//   - rows must be a slice of pointers.
//   - an individual row must have an ID field of [uuid.UUID].
func NewMetadata[T any](rows []T) Metadata {
	var metadata Metadata
	length := len(rows)
	metadata.Length = int64(length)

	if length < 1 {
		return metadata
	}

	lastRow := rows[length-1]
	deref := reflect.Indirect(reflect.ValueOf(lastRow))
	if !deref.IsValid() {
		return metadata
	}

	field := deref.FieldByName("ID")
	if !field.IsValid() || !field.CanInterface() {
		return metadata
	}
	if id, ok := field.Interface().(uuid.UUID); ok {
		metadata.LastSeen = id
	}

	return metadata
}
