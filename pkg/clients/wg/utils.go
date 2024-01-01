package wg

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func JoinInt64(values []int64, sep string) string {
	return strings.Join(lo.Map(values, func(i int64, _ int) string {
		return strconv.FormatInt(i, 10)
	}), sep)
}
