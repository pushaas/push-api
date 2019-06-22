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
		config *viper.Viper
		logger *zap.Logger
		redisClient redis.UniversalClient
	}
)

const (
	PublishingSuccess PublishingResult = iota
	PublishingInvalid
	PublishingFailure
)

func (s *publicationService) PublishMessage(message *models.Message) PublishingResult {
	channel := s.config.GetString("redis.pubsub.messages")

	bytes, err := json.Marshal(message)
	if err != nil {
		s.logger.Error("error marshaling message", zap.Any("message", message), zap.Error(err))
		return PublishingInvalid
	}
	messageJson := string(bytes)

	// TODO analyze whether val can be greater than 1
	val, err := s.redisClient.Publish(channel, messageJson).Result()

	if err == redis.Nil || err != nil {
		s.logger.Error("error publishing message", zap.Any("message", message), zap.Error(err))
		return PublishingFailure
	}

	s.logger.Debug("message published", zap.Int64("val", val), zap.Any("message", message))
	return PublishingSuccess
}

func NewPublicationService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) PublicationService {
	return &publicationService{
		config: config,
		logger: logger.Named("publicationService"),
		redisClient: redisClient,
	}
}
