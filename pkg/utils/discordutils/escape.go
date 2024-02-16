package du

import "strings"

var escapeReplacements = [][2]string{
	{"\\", "\\\\"},
	{"*", "\\*"},
	{"_", "\\_"},
	{"~", "\\~"},
	{"`", "\\`"},
	{"|", "\\|"},
	{"@", "@\u200b"},
}

func EscapeText(msg string) string {
	for _, r := range escapeReplacements {
		msg = strings.ReplaceAll(msg, r[0], r[1])
	}

	return msg
}
