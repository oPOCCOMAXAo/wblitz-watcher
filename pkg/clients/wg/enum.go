package wg

type App string

const (
	AppWotBlitz App = "wotb"
)

type Method string

const (
	MethodAccountInfo Method = "account/info"
	MethodAccountList Method = "account/list"
	MethodClansInfo   Method = "clans/info"
	MethodClansList   Method = "clans/list"
)
