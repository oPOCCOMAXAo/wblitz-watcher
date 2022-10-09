package repo

import "wblitz-watcher/wg/types"

type ClanID struct {
	ID     int    `bson:"id"`
	Region string `bson:"region"`
}

type Clan struct {
	ID       ClanID `bson:"_id"`
	Name     string `bson:"name"`
	Tag      string `bson:"tag"`
	Members  []int  `bson:"members"`
	LeaderID int    `bson:"leaderID"`
}

func (Clan) FromType(clan *types.ClanInfo) *Clan {
	return &Clan{
		ID: ClanID{
			ID:     clan.ClanID,
			Region: clan.Region,
		},
		Name:     clan.Name,
		Tag:      clan.Tag,
		Members:  clan.MembersIDs,
		LeaderID: clan.LeaderID,
	}
}

func (Clan) ToType(clan *Clan) *types.ClanInfo {
	return &types.ClanInfo{
		ClanID:       clan.ID.ID,
		Name:         clan.Name,
		Tag:          clan.Tag,
		MembersCount: len(clan.Members),
		MembersIDs:   clan.Members,
		LeaderID:     clan.LeaderID,
		Region:       clan.ID.Region,
	}
}
