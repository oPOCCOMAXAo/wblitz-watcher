package models

type WGClanMember struct {
	Region    Region
	ClanID    int64
	AccountID int64
}

type WGClanMembers struct {
	ID        WGClanID
	MembersID []int64
}
