package types

import (
	"strconv"
	"strings"
)

type MaybeInt int

func (i *MaybeInt) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	value, err := strconv.Atoi(s)
	if err == nil {
		*i = MaybeInt(value)
	}

	return err
}
