package models

type WGClanID struct {
	ID     int64
	Region Region
}

type WGClan struct {
	ID               int64
	Region           Region
	Tag              string
	Name             string
	UpdatedAt        int64
	MembersUpdatedAt int64
}
