package models

import "strings"

type Region string

const (
	RegionRU   Region = "ru"
	RegionEU   Region = "eu"
	RegionNA   Region = "na"
	RegionAsia Region = "asia"
)

func (r Region) Pretty() string {
	return strings.ToUpper(string(r))
}
