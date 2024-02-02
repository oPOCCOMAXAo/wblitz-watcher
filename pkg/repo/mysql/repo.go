package mysql

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/strings"
)

var _ do.Shutdownable = (*Repository)(nil)

type Repository struct {
	db *sql.DB
}

func New(
	db *sql.DB,
) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Shutdown() error {
	err := r.db.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repository) placeholders(count int) string {
	return strings.CopyJoin("?", ",", count)
}

func (r *Repository) placeholdersGroup(groupsCount, groupSize int) string {
	return strings.CopyJoin("("+r.placeholders(groupSize)+")", ",", groupsCount)
}
