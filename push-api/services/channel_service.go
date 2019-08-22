package services

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/pushaas/push-api/push-api/models"
)

type (
	ChannelCreationResult  int
	ChannelRetrievalResult int
	ChannelDeletionResult  int

	ChannelService interface {
		Create(channel *models.Channel) ChannelCreationResult
		Get(id string) (*models.Channel, ChannelRetrievalResult)
		GetAll() ([]*models.Channel, ChannelRetrievalResult)
		Delete(id string) ChannelDeletionResult
	}

	channelService struct{
		channelKeyPrefix string
		logger *zap.Logger
		redisClient redis.UniversalClient
	}
)

const (
	ChannelCreationSuccess ChannelCreationResult = iota
	ChannelCreationAlreadyExist
	ChannelCreationFailure
)

const (
	ChannelRetrievalSuccess ChannelRetrievalResult = iota
	ChannelRetrievalNotFound
	ChannelRetrievalFailure
)

const (
	ChannelDeletionSuccess ChannelDeletionResult = iota
	ChannelDeletionNotFound
	ChannelDeletionFailure
)

func (s *channelService) channelKey(suffix string) string {
	return fmt.Sprintf("%s_%s", s.channelKeyPrefix, suffix)
}

func (s *channelService) Create(channel *models.Channel) ChannelCreationResult {
	_, result := s.Get(channel.Id)

	if result == ChannelRetrievalSuccess {
		return ChannelCreationAlreadyExist
	}

	channel.Created = time.Now().UTC()

	key := s.channelKey(channel.Id)
	value, err := json.Marshal(channel)
	if err != nil {
		s.logger.Error("error while marshaling channel", zap.String("key", key), zap.String("channel", channel.Id), zap.Error(err))
		return ChannelCreationFailure
	}

	expiration := channel.Ttl * time.Second
	err = s.redisClient.Set(key, value, expiration).Err()
	if err != nil {
		s.logger.Error("error while saving channel", zap.String("key", key), zap.String("channel", channel.Id), zap.Error(err))
		return ChannelCreationFailure
	}

	s.logger.Debug("created channel", zap.Any("channel", channel))
	return ChannelCreationSuccess
}

func (s *channelService) Get(id string) (*models.Channel, ChannelRetrievalResult) {
	key := s.channelKey(id)

	value, err := s.redisClient.Get(key).Result()
	if err == redis.Nil {
		return nil, ChannelRetrievalNotFound
	} else if err != nil {
		return nil, ChannelRetrievalFailure
	}

	var channel models.Channel
	err = json.Unmarshal([]byte(value), &channel)
	if err != nil {
		return nil, ChannelRetrievalFailure
	}

	s.logger.Debug("retrieved channel", zap.Any("channel", channel))
	return &channel, ChannelRetrievalSuccess
}

func (s *channelService) GetAll() ([]*models.Channel, ChannelRetrievalResult) {
	key := s.channelKey("*")

	var keys []string
	iterator := s.redisClient.Scan(0, key, 10).Iterator()
	for iterator.Next() {
		keys = append(keys, iterator.Val())
	}

	if err := iterator.Err(); err != nil {
		s.logger.Error("error while retrieving iterating all channels", zap.Error(err))
		return nil, ChannelRetrievalFailure
	}

	if len(keys) == 0 {
		return make([]*models.Channel, 0), ChannelRetrievalSuccess
	}

	results, err := s.redisClient.MGet(keys...).Result()
	if err != nil {
		s.logger.Error("error while retrieving all channels", zap.Error(err))
		return nil, ChannelRetrievalFailure
	}

	channels := make([]*models.Channel, len(results))
	for i, v := range results {
		var channel models.Channel
		if strValue, ok := v.(string); ok {
			err = json.Unmarshal([]byte(strValue), &channel)
			if err != nil {
				s.logger.Error("error while unmarshalling all channels", zap.String("value", strValue), zap.Error(err))
				return nil, ChannelRetrievalFailure
			}
			channels[i] = &channel
		}
	}

	sort.Slice(channels, func(i, j int) bool {
		return channels[i].Created.String() < channels[j].Created.String()
	})

	return channels, ChannelRetrievalSuccess
}

func (s *channelService) Delete(id string) ChannelDeletionResult {
	key := s.channelKey(id)
	result, err := s.redisClient.Del(key).Result()

	if err == redis.Nil {
		return ChannelDeletionNotFound
	} else if err != nil {
		return ChannelDeletionFailure
	}

	if result == 0 {
		return ChannelDeletionNotFound
	}

	s.logger.Debug("deleted channel", zap.String("id", id))
	return ChannelDeletionSuccess
}

func NewChannelService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) ChannelService {
	channelKeyPrefix := config.GetString("redis.db.channel.prefix")

	return &channelService{
		channelKeyPrefix: channelKeyPrefix,
		logger: logger.Named("channelService"),
		redisClient: redisClient,
	}
}

