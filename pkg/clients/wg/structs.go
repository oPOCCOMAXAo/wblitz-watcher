package wg

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

type Request struct {
	Region models.Region
	App    App
	Method Method
	Data   url.Values
}

type ResponseError struct {
	Code    int    `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

func (e *ResponseError) Text() string {
	lines := []string{
		strconv.Itoa(e.Code) + ": " + e.Message,
	}

	if e.Field != "" || e.Value != "" {
		lines = append(lines, e.Field+" "+e.Value)
	}

	return strings.Join(lines, "\n")
}

func (e *ResponseError) GetError() error {
	return fmt.Errorf("%w %s", ErrAPI, e.Text())
}

type ResponseMeta struct {
	Count int64 `json:"count"`
}

type Response struct {
	Status string         `json:"status"`
	Error  *ResponseError `json:"error"`
	Meta   ResponseMeta   `json:"meta"`
	Data   any            `json:"data"`
}

type ClanInfo struct {
	MembersCount int64         `json:"members_count"`
	Name         string        `json:"name"`
	CreatorName  string        `json:"creator_name"`
	ClanID       int64         `json:"clan_id"`
	LeaderID     int64         `json:"leader_id"`
	MembersIDs   []int64       `json:"members_ids"`
	Tag          string        `json:"tag"`
	Region       models.Region `json:"-"`
}

func (c *ClanInfo) WGClanID() models.WGClanID {
	return models.WGClanID{
		ID:     c.ClanID,
		Region: c.Region,
	}
}

func (c *ClanInfo) StatName() string {
	return fmt.Sprintf("[**%s**] %s / %s", c.Tag, c.Region.Pretty(), c.Name)
}

type StatisticsEntry struct {
	Spotted              int64 `json:"spotted"`
	MaxFragsTankID       int64 `json:"max_frags_tank_id"`
	Hits                 int64 `json:"hits"`
	Frags                int64 `json:"frags"`
	MaxXP                int64 `json:"max_xp"`
	MaxXPTankID          int64 `json:"max_xp_tank_id"`
	Wins                 int64 `json:"wins"`
	Losses               int64 `json:"losses"`
	CapturePoints        int64 `json:"capture_points"`
	Battles              int64 `json:"battles"`
	DamageDealt          int64 `json:"damage_dealt"`
	DamageReceived       int64 `json:"damage_received"`
	MaxFrags             int64 `json:"max_frags"`
	Shots                int64 `json:"shots"`
	Frags8p              int64 `json:"frags8p"`
	XP                   int64 `json:"xp"`
	WinAndSurvived       int64 `json:"win_and_survived"`
	SurvivedBattles      int64 `json:"survived_battles"`
	DroppedCapturePoints int64 `json:"dropped_capture_points"`
}

type AccountStatistics struct {
	Clan StatisticsEntry `json:"clan"`
	All  StatisticsEntry `json:"all"`
}

type AccountInfo struct {
	AccountID      int64             `json:"account_id"`
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

type AccountListEntry struct {
	AccountID int64  `json:"account_id"`
	Nickname  string `json:"nickname"`
}

type ClanListEntry struct {
	ClanID       int64  `json:"clan_id"`
	Tag          string `json:"tag"`
	Name         string `json:"name"`
	MembersCount int64  `json:"members_count"`
}
