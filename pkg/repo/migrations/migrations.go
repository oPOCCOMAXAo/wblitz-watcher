package migrations

import (
	"database/sql"
	"embed"
	"errors"
	"io/fs"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

type Migration struct {
	ID       string
	Migrate  func(*sql.DB) error
	Rollback func(*sql.DB) error
}

var _ FS = embed.FS{}

type FS interface {
	fs.ReadDirFS
	fs.ReadFileFS
}

func GetMigrations(fsys FS) ([]*Migration, error) {
	ids, err := listID(fsys)
	if err != nil {
		return nil, err
	}

	res := []*Migration{}

	for _, id := range ids {
		migration, err := loadMigration(fsys, id)
		if err != nil {
			return nil, err
		}

		res = append(res, migration)
	}

	return res, nil
}

func listID(fsys FS) ([]string, error) {
	dir, err := fsys.ReadDir("sql")
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	ids := map[string]struct{}{}

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		name := strings.Split(entry.Name(), ".")

		const totalSegments = 3
		if len(name) < totalSegments {
			continue
		}

		if name[len(name)-1] != "sql" {
			continue
		}

		ids[name[0]] = struct{}{}
	}

	return maps.Keys(ids), nil
}

func loadMigration(fsys FS, id string) (*Migration, error) {
	res := Migration{
		ID: id,
	}

	data, err := readFile(fsys, "sql/"+id+".up.sql")
	if err != nil {
		return nil, err
	}

	if data != nil {
		res.Migrate = wrapRawSQL(data)
	}

	data, err = readFile(fsys, "sql/"+id+".down.sql")
	if err != nil {
		return nil, err
	}

	if data != nil {
		res.Rollback = wrapRawSQL(data)
	}

	return &res, nil
}

func readFile(fsys FS, file string) ([]byte, error) {
	data, err := fsys.ReadFile(file)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		//nolint:wrapcheck
		return nil, err
	}

	return data, nil
}

func prepareStatements(rawData []byte) []string {
	sqls := strings.Split(string(rawData), ";")

	return lo.FilterMap(sqls, func(sql string, _ int) (string, bool) {
		sql = strings.Trim(sql, "\n\r \t")

		return sql, len(sql) != 0
	})
}

func wrapRawSQL(rawData []byte) func(*sql.DB) error {
	return func(db *sql.DB) error {
		sqls := prepareStatements(rawData)

		for _, sql := range sqls {
			_, err := db.Exec(sql)
			if err != nil {
				//nolint:wrapcheck
				return err
			}
		}

		return nil
	}
}
