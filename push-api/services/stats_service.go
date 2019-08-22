package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/pushaas/push-api/push-api/models"
)

type (
	StatsRetrievalResult int

	StatsService interface {
		GetChannelStats(string) (*models.ChannelStatsResult, StatsRetrievalResult)
		GetGlobalStats() (*models.GlobalStatsResult, StatsRetrievalResult)
	}

	statsService struct{
		channelStatsPrefix string
		globalStatsPrefix string
		logger            *zap.Logger
		redisClient       redis.UniversalClient
	}
)

const (
	StatsRetrievalSuccess StatsRetrievalResult = iota
	StatsRetrievalNotFound
	StatsRetrievalFailure
)

func (s *statsService) channelStatsKey(suffix string) string {
	return fmt.Sprintf("%s:%s", s.channelStatsPrefix, suffix)
}

func (s *statsService) globalStatsKey(suffix string) string {
	return fmt.Sprintf("%s:%s", s.globalStatsPrefix, suffix)
}

func (s *statsService) buildChannelStatsResult(stats []models.ChannelStats) *models.ChannelStatsResult {
	// aggregate
	aggregated := &models.ChannelStatsAggregated{}
	for _, stat := range stats {
		aggregated.Subscribers += stat.Subscribers
	}

	return &models.ChannelStatsResult{
		Aggregated: aggregated,
		All: stats,
	}
}

func (s *statsService) buildGlobalStatsResult(stats []models.GlobalStats) *models.GlobalStatsResult {
	// aggregate
	aggregated := &models.GlobalStatsAggregated{}
	for _, stat := range stats {
		if aggregated.Time == "" {
			aggregated.Time = stat.Time
		}
		if stat.Time < aggregated.Time {
			aggregated.Time = stat.Time
		}
		aggregated.Subscribers += stat.Subscribers
	}

	return &models.GlobalStatsResult{
		Aggregated: aggregated,
		All: stats,
	}
}


func (s *statsService) getStatsData(key, description string) ([]interface{}, StatsRetrievalResult) {
	// get keys
	keys, err := s.redisClient.Keys(key).Result()
	if err != nil {
		s.logger.Error(fmt.Sprintf("failed to retrieve %s stats keys", description), zap.String("key", key), zap.Error(err))
		return nil, StatsRetrievalFailure
	}

	if len(keys) == 0 {
		return nil, StatsRetrievalNotFound
	}

	// get data
	strStats, err := s.redisClient.MGet(keys...).Result()
	if err != nil {
		s.logger.Error(fmt.Sprintf("failed to retrieve %s stats values", description), zap.String("key", key), zap.Error(err))
		return nil, StatsRetrievalFailure
	}

	return strStats, StatsRetrievalSuccess
}

func (s *statsService) GetChannelStats(channel string) (*models.ChannelStatsResult, StatsRetrievalResult) {
	key := s.channelStatsKey(fmt.Sprintf("%s:*", channel))
	strStats, result := s.getStatsData(key, fmt.Sprintf("channel %s", channel))
	if result != StatsRetrievalSuccess {
		return nil, result
	}

	// transform individual data
	stats := make([]models.ChannelStats, len(strStats))
	for i, strStat := range strStats {
		str := strStat.(string)
		var stat models.ChannelStats
		_ = json.Unmarshal([]byte(str), &stat)
		stats[i] = stat
	}

	return s.buildChannelStatsResult(stats), StatsRetrievalSuccess
}

func (s *statsService) GetGlobalStats() (*models.GlobalStatsResult, StatsRetrievalResult) {
	key := s.globalStatsKey("*")
	strStats, result := s.getStatsData(key, "global")
	if result != StatsRetrievalSuccess {
		return nil, result
	}

	// transform individual data
	stats := make([]models.GlobalStats, len(strStats))
	for i, strStat := range strStats {
		str := strStat.(string)
		var stat models.GlobalStats
		_ = json.Unmarshal([]byte(str), &stat)
		stats[i] = stat
	}

	return s.buildGlobalStatsResult(stats), StatsRetrievalSuccess
}

func NewStatsService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) StatsService {
	channelStatsPrefix := config.GetString("redis.db.stats_channel.prefix")
	globalStatsPrefix := config.GetString("redis.db.stats_global.prefix")

	return &statsService{
		channelStatsPrefix: channelStatsPrefix,
		globalStatsPrefix: globalStatsPrefix,
		logger:            logger.Named("statsService"),
		redisClient:       redisClient,
	}
}
