package services

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/models"
)

type (
	MessagePublishingResult int

	PublicationService interface {
		PublishMessage(*models.Message) MessagePublishingResult
	}

	publicationService struct{
		config *viper.Viper
		logger *zap.Logger
		redisClient redis.UniversalClient
	}
)

const (
	MessagePublishingSuccess MessagePublishingResult = iota
	MessagePublishingFailure
	// TODO maybe add a MessagePublishingChannelNotFound
)

// TODO support array of channel aids on `message`
func (s *publicationService) PublishMessage(message *models.Message) MessagePublishingResult {
	channel := s.config.GetString("redis.channels.messages") // TODO probably we won't have this single channel

	// TODO analyze whether val can be greater than 1
	_, err := s.redisClient.Publish(channel, message.Content).Result()

	if err == redis.Nil || err != nil {
		s.logger.Error("error publishing message", zap.Any("message", message), zap.Error(err))
		return MessagePublishingFailure
	}

	s.logger.Debug("message published", zap.Any("message", message))
	return MessagePublishingSuccess
}

func NewPublicationService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) PublicationService {
	return &publicationService{
		config: config,
		logger: logger.Named("publicationService"),
		redisClient: redisClient,
	}
}
