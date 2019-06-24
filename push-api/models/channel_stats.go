package models

type (
	ChannelStats struct {
		Channel           string `json:"channel"`
		PublishedMessages int64  `json:"published_messages"`
		StoredMessages    int64  `json:"stored_messages"`
		Subscribers       int64  `json:"subscribers"`
		Hostname          string `json:"hostname"`
		Agent             string `json:"agent"`
	}

	ChannelStatsAggregated struct {
		Subscribers int64 `json:"subscribers"`
	}

	ChannelStatsResult struct {
		Aggregated *ChannelStatsAggregated `json:"aggregated"`
		All        []ChannelStats          `json:"all"`
	}
)
