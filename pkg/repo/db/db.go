package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func OpenMySQL(
	dsn string,
) (*sql.DB, error) {
	res, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.SetMaxOpenConns(10)
	res.SetMaxIdleConns(5)
	res.SetConnMaxIdleTime(time.Minute)
	res.SetConnMaxLifetime(time.Minute * 5)

	err = res.Ping()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
