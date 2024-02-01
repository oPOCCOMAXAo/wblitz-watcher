package models

import "fmt"

type WGClanMember struct {
	Region    Region
	ClanID    int64
	AccountID int64
}

func (m *WGClanMember) GetFullClanID() WGClanID {
	return WGClanID{
		Region: m.Region,
		ID:     m.ClanID,
	}
}

type WGClanMembers struct {
	ID         WGClanID
	MembersIDs []int64
	IsFound    bool // api sometimes returns null instead of object. this flag is for that case.
}

func (c *WGClanMembers) EntityUniqueID() string {
	return fmt.Sprintf("clan#%s#%d", c.ID.Region.Pretty(), c.ID.ID)
}
