package watcher

type Color int

func ColorRGB(r, g, b byte) Color {
	return Color(r)<<16 + Color(g)<<8 + Color(b)
}

const (
	ColorError    Color = 0xff0000
	ColorEnter    Color = 0x7fff7f
	ColorLeave    Color = 0xff7f7f
	ColorWin      Color = 0x006600
	ColorLoss     Color = 0x660000
	ColorDraw     Color = 0x666600
	ColorDisabled Color = 0x010000
	ColorEnabled  Color = 0xfefffe
)

const ClanInitializationIntervalSeconds = 24 * 60 * 60

// messages.
const (
	MessageError   = "Error"
	MessageEnter   = "Enter"
	MessageLeave   = "Leave"
	MessageWins    = "Wins"
	MessageDamage  = "Damage"
	MessageBattles = "Battles"
	MessageTag     = "Tag"
	MessageName    = "Name"
	MessageRegion  = "Region"
)
