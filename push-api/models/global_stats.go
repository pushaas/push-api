package models

type (
	GlobalStats struct {
		Hostname          string `json:"hostname"`
		Time              string `json:"time"`
		Channels          int64  `json:"channels"`
		WildcardChannels  int64  `json:"wildcard_channels"`
		PublishedMessages int64  `json:"published_messages"`
		StoredMessages    int64  `json:"stored_messages"`
		MessagesInTrash   int64  `json:"messages_in_trash"`
		ChannelsInDelete  int64  `json:"channels_in_delete"`
		ChannelsInTrash   int64  `json:"channels_in_trash"`
		Subscribers       int64  `json:"subscribers"`
		Uptime            int64  `json:"uptime"`
		Agent             string `json:"agent"`
	}

	GlobalStatsAggregated struct {
		Time        string `json:"time"`
		Subscribers int64  `json:"subscribers"`
	}

	GlobalStatsResult struct {
		Aggregated *GlobalStatsAggregated `json:"aggregated"`
		All        []GlobalStats          `json:"all"`
	}
)
