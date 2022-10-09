package types

import (
	"fmt"
)

type ClanInfo struct {
	MembersCount int    `json:"members_count"`
	Name         string `json:"name"`
	CreatorName  string `json:"creator_name"`
	ClanID       int    `json:"clan_id"`
	LeaderID     int    `json:"leader_id"`
	MembersIDs   []int  `json:"members_ids"`
	Tag          string `json:"tag"`
	Region       string `json:"-"`
}

func (c *ClanInfo) StatName() string {
	return fmt.Sprintf("[**%s**] %s", c.Tag, c.Name)
}