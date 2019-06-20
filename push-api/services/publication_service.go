package services

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	PublicationService interface {
		Publish() error
	}

	publicationService struct{
		config *viper.Viper
		logger *zap.Logger
		redisClient redis.UniversalClient
	}
)

func (s *publicationService) Publish() error {
	channel := s.config.GetString("redis.channels.messages")
	// TODO
	s.redisClient.Publish(channel, "meu nome é rafael")
	return nil
}

func NewPublicationService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) PublicationService {
	return &publicationService{
		config: config,
		logger: logger.Named("publicationService"),
		redisClient: redisClient,
	}
}
