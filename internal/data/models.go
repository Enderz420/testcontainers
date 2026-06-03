package data

import (
	"database/sql"
	"time"
)

type Models struct {
	Users       UserModel
	Blogpost    BlogpostModel
	Groups      GroupModel
	Permissions PermissionModel
}

func NewModels(db *sql.DB, timeout *time.Duration) Models {
	return Models{
		Users:       UserModel{DB: db, Timeout: timeout},
		Blogpost:    BlogpostModel{DB: db, Timeout: timeout},
		Groups:      GroupModel{DB: db, Timeout: timeout},
		Permissions: PermissionModel{DB: db, Timeout: timeout},
	}
}
