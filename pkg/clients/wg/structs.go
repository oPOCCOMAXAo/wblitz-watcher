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
	Conunt int `json:"count"`
}

type Response struct {
	Status string         `json:"status"`
	Error  *ResponseError `json:"error"`
	Meta   ResponseMeta   `json:"meta"`
	Data   any            `json:"data"`
}

type ClanInfo struct {
	MembersCount int           `json:"members_count"`
	Name         string        `json:"name"`
	CreatorName  string        `json:"creator_name"`
	ClanID       int           `json:"clan_id"`
	LeaderID     int           `json:"leader_id"`
	MembersIDs   []int         `json:"members_ids"`
	Tag          string        `json:"tag"`
	Region       models.Region `json:"-"`
}

func (c *ClanInfo) StatName() string {
	return fmt.Sprintf("[**%s**] %s / %s", c.Tag, c.Region.Pretty(), c.Name)
}

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

type AccountListEntry struct {
	AccountID int64  `json:"account_id"`
	Nickname  string `json:"nickname"`
}

type ClanListEntry struct {
	ClanID       int64  `json:"clan_id"`
	Tag          string `json:"tag"`
	Name         string `json:"name"`
	MembersCount int    `json:"members_count"`
}
