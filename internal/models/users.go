package models

import (
	"database/sql"
	"time"

	mssql "github.com/microsoft/go-mssqldb"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID        mssql.UniqueIdentifier `json:"id"`
	Username  string                 `json:"username"`
	Email     string                 `json:"email"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

func (m *UserModel) selectOne(id mssql.UniqueIdentifier) (*User, error) {
	const stmt string = `SELECT ID, Username, Email, CreatedAt, UpdatedAt FROM Users WHERE ID = ?`

	var user User

	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
