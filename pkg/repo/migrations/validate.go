package migrations

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
)

const MigrationPrefixFormat = "200601021504"

var regexpIDPrefix = regexp.MustCompile(`^\d{12}`)

func validate(migrations []*Migration) error {
	for _, migration := range migrations {
		err := validateID(migration.ID)
		if err != nil {
			return errors.WithMessagef(err, "migration#%s has invalid id %+v", migration.ID, err)
		}

		if migration.Migrate == nil {
			return errors.WithMessagef(err, "migration#%s must have Migrate", migration.ID)
		}
	}

	return nil
}

func validateID(id string) error {
	if len(id) < len(MigrationPrefixFormat) {
		return errors.WithStack(ErrTooShortID)
	}

	if !regexpIDPrefix.MatchString(id) {
		return errors.WithStack(ErrInvalidTimeFormat) // must be 10 digits.
	}

	id = id[:len(MigrationPrefixFormat)]

	idDate, err := time.ParseInLocation(MigrationPrefixFormat, id, time.Local)
	if err != nil {
		return errors.WithStack(err) // must be parseable by time.
	}

	now := time.Now()
	if now.Before(idDate) {
		return errors.WithStack(ErrInvalidDate) // must be in past.
	}

	dateString := idDate.Format(MigrationPrefixFormat)
	if dateString != id {
		return errors.WithStack(ErrInvalidTimeFormat) // must be in valid form.
	}

	return nil
}
