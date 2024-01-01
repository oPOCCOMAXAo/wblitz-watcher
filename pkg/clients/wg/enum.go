package wg

import (
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/maps"
)

type App string

const (
	AppWotBlitz App = "wotb"
)

type Region uint8

const (
	RegionRU   Region = 0
	RegionEU   Region = 1
	RegionNA   Region = 2
	RegionAsia Region = 3

	RegionUnknown Region = 255
)

func (r Region) Host() string {
	return maps.GetDefault(map[Region]string{
		RegionAsia:    "api.wotblitz.asia",
		RegionEU:      "api.wotblitz.eu",
		RegionNA:      "api.wotblitz.com",
		RegionRU:      "api.tanki.su",
		RegionUnknown: "api.wotblitz.eu",
	}, r, RegionUnknown)
}

func (r Region) Name() string {
	return maps.GetDefault(map[Region]string{
		RegionAsia:    "ASIA",
		RegionEU:      "EU",
		RegionNA:      "NA",
		RegionRU:      "RU",
		RegionUnknown: "EU",
	}, r, RegionUnknown)
}

func RegionFromName(name string) Region {
	return maps.GetDefault(map[string]Region{
		"ASIA": RegionAsia,
		"EU":   RegionEU,
		"NA":   RegionNA,
		"RU":   RegionRU,
		"":     RegionUnknown,
	}, name, "")
}

type Method string

const (
	MethodClansInfo   Method = "clans/info"
	MethodAccoutInfo  Method = "account/info"
	MethodAccountList Method = "account/list"
)
