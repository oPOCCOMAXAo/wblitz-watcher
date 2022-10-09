package app

import (
	"wblitz-watcher/app/diff"
)

const (
	AppTag = "wblitz"
)

const (
	DiffCreated diff.Type = iota
	DiffName
	DiffTag
	DiffLeader
	DiffEnter
	DiffLeave
)

const (
	StringEnter = "Вступление в клан"
	StringLeave = "Выход из клана"
	StringWin   = "ПОБЕДА"
	StringLoss  = "ПОРАЖЕНИЕ"
	StringDraw  = "НИЧЬЯ"
)
