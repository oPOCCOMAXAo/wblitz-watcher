package jsonutils

import (
	"strconv"
	"strings"
)

type MaybeInt int64

func (i *MaybeInt) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	value, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*i = MaybeInt(value)
	}

	//nolint:wrapcheck
	return err
}
