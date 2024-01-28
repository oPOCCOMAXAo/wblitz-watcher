package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // for MySQL driver.
	"github.com/pkg/errors"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

func OpenMySQL(
	ctx context.Context,
	dsn string,
) (*sql.DB, error) {
	res, err := otelsql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res.SetMaxOpenConns(10)
	res.SetMaxIdleConns(5)
	res.SetConnMaxIdleTime(time.Minute)
	res.SetConnMaxLifetime(time.Minute * 5)

	err = res.PingContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
