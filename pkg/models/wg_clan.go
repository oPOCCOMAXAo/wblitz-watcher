package models

import "fmt"

type WGClanID struct {
	ID     int64
	Region Region
}

func (id WGClanID) MapMember(memberID int64, _ int) *WGClanMember {
	return &WGClanMember{
		Region:    id.Region,
		ClanID:    id.ID,
		AccountID: memberID,
	}
}

type WGClan struct {
	ID               int64
	Region           Region
	Tag              string
	Name             string
	UpdatedAt        int64
	MembersUpdatedAt int64
}

func (c *WGClan) EntityUniqueID() string {
	return fmt.Sprintf("clan#%s#%d", c.Region.Pretty(), c.ID)
}

func (c *WGClan) GetFullClanID() WGClanID {
	return WGClanID{
		ID:     c.ID,
		Region: c.Region,
	}
}

type WGClanExtended struct {
	Clan       *WGClan
	MembersIDs []int64
}
