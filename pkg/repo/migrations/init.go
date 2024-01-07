package migrations

import (
	"database/sql"
	"embed"
	"sort"

	"github.com/pkg/errors"
)

//go:embed sql/*.sql
var sqlDir embed.FS

func Init(db *sql.DB) error {
	migrations, err := GetMigrations(sqlDir)
	if err != nil {
		return err
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].ID < migrations[j].ID
	})

	err = validate(migrations)
	if err != nil {
		return err
	}

	ids, err := initMigrator(db)
	if err != nil {
		return err
	}

	err = applyMigrations(db, ids, migrations)
	if err != nil {
		return err
	}

	return nil
}

func initMigrator(
	db *sql.DB,
) (map[string]struct{}, error) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS migrations (id VARCHAR(255) PRIMARY KEY)`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := db.Query(`SELECT id FROM migrations`)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.WithStack(err)
	}

	res := map[string]struct{}{}

	for rows.Next() {
		var id string

		err = rows.Scan(&id)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		res[id] = struct{}{}
	}

	return res, nil
}

func applyMigrations(
	db *sql.DB,
	ids map[string]struct{},
	migrations []*Migration,
) error {
	for _, migration := range migrations {
		if _, ok := ids[migration.ID]; ok {
			continue
		}

		err := applyMigration(db, migration)
		if err != nil {
			return err
		}

		ids[migration.ID] = struct{}{}
	}

	return nil
}

func applyMigration(
	db *sql.DB,
	migration *Migration,
) (err error) {
	defer func() {
		if err != nil {
			err = errors.WithMessagef(err, "migration#%s failed", migration.ID)

			_, _ = db.Exec(`DELETE FROM migrations WHERE id = ?`, migration.ID)

			if migration.Rollback != nil {
				_ = migration.Rollback(db)
			}
		}
	}()

	err = migration.Migrate(db)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = db.Exec(`INSERT INTO migrations (id) VALUES (?)`, migration.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
