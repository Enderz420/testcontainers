package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

type Permission struct {
	ID           mssql.UniqueIdentifier `json:"id"`
	GroupID      mssql.UniqueIdentifier `json:"group_id"`
	Name         string                 `json:"name"`
	CreatedAt    time.Time              `json:"created_at"`
	LastModified time.Time              `json:"last_modified"`
}

type PermissionModel struct {
	DB      *sql.DB
	Timeout *time.Duration
}

func (m *PermissionModel) SelectOne(ctx context.Context, id uuid.UUID) (*Permission, error) {
	ctx, cancel := context.WithTimeout(ctx, *m.Timeout)
	defer cancel()

	const stmt = `
	SELECT
		id,
		group_id,
		name,
		created_at,
		last_modified
	FROM core.permissions
	WHERE id = @ID`

	row := m.DB.QueryRowContext(ctx, stmt, sql.Named("ID", id))

	var p Permission
	if err := row.Scan(&p.ID, &p.GroupID, &p.Name, &p.CreatedAt, &p.LastModified); err != nil {
		return nil, err
	}

	return &p, nil
}
