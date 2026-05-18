package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func OpenDB(config Config) (*sql.DB, error) {
	fmt.Println("Using:", config.DSN)
	db, err := sql.Open("sqlserver", config.DSN)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)

	duration, err := time.ParseDuration(config.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}