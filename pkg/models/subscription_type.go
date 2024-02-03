package models

type SubscriptionType string

const (
	STClan SubscriptionType = "clan"
	STInfo SubscriptionType = "info"
)

//nolint:gochecknoglobals
var SubscriptionTypes = []SubscriptionType{
	STClan,
	STInfo,
}
