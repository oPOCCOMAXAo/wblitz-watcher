package models

type SubscriptionClan struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement"`
	InstanceID int64  `gorm:"column:instance_id"`
	ClanID     int64  `gorm:"column:clan_id"`
	Region     Region `gorm:"column:region;enum:ru,eu,na,asia"`
}

func (SubscriptionClan) TableName() string {
	return "subscription_clan"
}
