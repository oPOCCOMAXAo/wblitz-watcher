package models

type BotInstance struct {
	ID        int64            `gorm:"column:id;primaryKey;autoIncrement"`
	ServerID  string           `gorm:"column:server_id;size:64;index"`
	ChannelID string           `gorm:"column:channel_id;size:64"`
	Type      SubscriptionType `gorm:"column:type;enum:clan"`
	CreatedAt int64            `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64            `gorm:"column:updated_at;autoUpdateTime"`
}

func (BotInstance) TableName() string {
	return "bot_instance"
}
