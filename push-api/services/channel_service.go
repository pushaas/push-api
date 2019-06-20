package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/models"
)

type (
	ChannelCreationResult  int
	ChannelRetrievalResult int
	ChannelDeletionResult  int

	ChannelService interface {
		Create(channel *models.Channel) ChannelCreationResult
		Get(id string) (*models.Channel, ChannelRetrievalResult)
		Delete(id string) ChannelDeletionResult
	}

	channelService struct{
		config *viper.Viper
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

const channelPrefix = "ch"

func channelKey(id string) string {
	return fmt.Sprintf("%s_%s", channelPrefix, id)
}

func (s *channelService) Create(channel *models.Channel) ChannelCreationResult {
	_, result := s.Get(channel.Id)

	if result == ChannelRetrievalSuccess {
		return ChannelCreationAlreadyExist
	}

	key := channelKey(channel.Id)
	value, err := json.Marshal(channel)
	if err != nil {
		s.logger.Error("error while marshaling channel", zap.String("key", key), zap.String("channel", channel.Id), zap.Error(err))
		return ChannelCreationFailure
	}

	err = s.redisClient.Set(key, value, channel.Ttl).Err()
	if err != nil {
		s.logger.Error("error while saving channel", zap.String("key", key), zap.String("channel", channel.Id), zap.Error(err))
		return ChannelCreationFailure
	}

	return ChannelCreationSuccess
}

func (s *channelService) Get(id string) (*models.Channel, ChannelRetrievalResult) {
	key := channelKey(id)

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

	return &channel, ChannelRetrievalSuccess
}

func (s *channelService) Delete(id string) ChannelDeletionResult {
	key := channelKey(id)
	result, err := s.redisClient.Del(key).Result()

	if err == redis.Nil {
		return ChannelDeletionNotFound
	} else if err != nil {
		return ChannelDeletionFailure
	}

	if result == 0 {
		return ChannelDeletionNotFound
	}

	return ChannelDeletionSuccess
}


func NewChannelService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) ChannelService {
	return &channelService{
		config: config,
		logger: logger.Named("channelService"),
		redisClient: redisClient,
	}
}

