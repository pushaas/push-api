package services

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/models"
)

type (
	PublishingResult int

	PublicationService interface {
		PublishMessage(*models.Message) PublishingResult
	}

	publicationService struct{
		logger *zap.Logger
		pubsubChannel string
		redisClient redis.UniversalClient
	}
)

const (
	PublishingSuccess PublishingResult = iota
	PublishingInvalid
	PublishingFailure
)

func (s *publicationService) PublishMessage(message *models.Message) PublishingResult {
	bytes, err := json.Marshal(message)
	if err != nil {
		s.logger.Error("error marshaling message", zap.Any("message", message), zap.Error(err))
		return PublishingInvalid
	}
	messageJson := string(bytes)

	channel := s.pubsubChannel
	numListeners, err := s.redisClient.Publish(channel, messageJson).Result()

	if err == redis.Nil || err != nil {
		s.logger.Error("error publishing message", zap.String("channel", channel), zap.Any("message", message), zap.Error(err))
		return PublishingFailure
	}

	s.logger.Debug("message published", zap.String("channel", channel), zap.Int64("numListeners", numListeners), zap.String("message", messageJson))
	return PublishingSuccess
}

func NewPublicationService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) PublicationService {
	pubsubChannel := config.GetString("redis.pubsub-channels.publish")

	return &publicationService{
		logger: logger.Named("publicationService"),
		pubsubChannel: pubsubChannel,
		redisClient: redisClient,
	}
}
