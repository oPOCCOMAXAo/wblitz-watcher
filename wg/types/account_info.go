package types

import (
	"fmt"
)

type StatisticsEntry struct {
	Spotted              int `json:"spotted"`
	MaxFragsTankID       int `json:"max_frags_tank_id"`
	Hits                 int `json:"hits"`
	Frags                int `json:"frags"`
	MaxXP                int `json:"max_xp"`
	MaxXPTankID          int `json:"max_xp_tank_id"`
	Wins                 int `json:"wins"`
	Losses               int `json:"losses"`
	CapturePoints        int `json:"capture_points"`
	Battles              int `json:"battles"`
	DamageDealt          int `json:"damage_dealt"`
	DamageReceived       int `json:"damage_received"`
	MaxFrags             int `json:"max_frags"`
	Shots                int `json:"shots"`
	Frags8p              int `json:"frags8p"`
	XP                   int `json:"xp"`
	WinAndSurvived       int `json:"win_and_survived"`
	SurvivedBattles      int `json:"survived_battles"`
	DroppedCapturePoints int `json:"dropped_capture_points"`
}

type AccountStatistics struct {
	Clan StatisticsEntry `json:"clan"`
	All  StatisticsEntry `json:"all"`
}

type AccountInfo struct {
	AccountID      int               `json:"account_id"`
	Nickname       string            `json:"nickname"`
	LastBattleTime int64             `json:"last_battle_time"`
	Statistics     AccountStatistics `json:"statistics"`
}

func (p *AccountInfo) AuthorName() string {
	return p.Nickname
}

func (p *AccountInfo) StatWins() string {
	if p.Statistics.All.Battles == 0 {
		return "0%"
	}

	return fmt.Sprintf("%.2f%%", float64(p.Statistics.All.Wins)/float64(p.Statistics.All.Battles)*100)
}

func (p *AccountInfo) StatDamage() string {
	if p.Statistics.All.Battles == 0 {
		return "0"
	}

	return fmt.Sprintf("%d", p.Statistics.All.DamageDealt/p.Statistics.All.Battles)
}

func (p *AccountInfo) StatBattles() string {
	return fmt.Sprintf("%d", p.Statistics.All.Battles)
}
