package api

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
)

func (r Region) Host() string {
	switch r {
	case RegionAsia:
		return "api.wotblitz.asia"
	case RegionEU:
		return "api.wotblitz.eu"
	case RegionNA:
		return "api.wotblitz.com"
	case RegionRU:
		return "api.wotblitz.ru"
	default:
		return "api.wotblitz.eu"
	}
}

func (r Region) Name() string {
	switch r {
	case RegionAsia:
		return "ASIA"
	case RegionEU:
		return "EU"
	case RegionNA:
		return "NA"
	case RegionRU:
		return "RU"
	default:
		return "EU"
	}
}

type Method string

const (
	MethodClansInfo Method = "clans/info"
)
