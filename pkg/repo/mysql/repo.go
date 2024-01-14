package mysql

import (
	"database/sql"

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
	//nolint:wrapcheck
	return r.db.Close()
}

func (r *Repository) placeholders(count int) string {
	return strings.CopyJoin("?", ",", count)
}

func (r *Repository) placeholdersGroup(groupsCount, groupSize int) string {
	return strings.CopyJoin("("+r.placeholders(groupSize)+")", ",", groupsCount)
}
