package discord

func (s *Service) isChannelIgnored(
	channelID string,
) bool {
	specifiedIgnore := s.ignoredChannelMap[channelID]
	notUsed := s.useOnlyChannels && !s.onlyChannelsMap[channelID]

	return specifiedIgnore || notUsed
}
