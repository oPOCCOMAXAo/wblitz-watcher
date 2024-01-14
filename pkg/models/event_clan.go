package models

type EventClan struct {
	ID          int64
	Time        int64
	Type        EventTypeClan
	Region      Region
	ClanID      int64
	AccountID   int64
	IsProcessed bool
}
