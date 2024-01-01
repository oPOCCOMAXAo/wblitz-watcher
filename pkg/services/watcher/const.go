package watcher

type Color int

func ColorRGB(r, g, b byte) Color {
	return Color(r)<<16 + Color(g)<<8 + Color(b)
}

const (
	ColorEnter Color = 0xfefffe
	ColorLeave Color = 0x010000
	ColorWin   Color = 0x006600
	ColorLoss  Color = 0x660000
	ColorDraw  Color = 0x666600
)
