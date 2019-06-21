package ctors

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/services"
)

func NewPublicationService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) services.PublicationService {
	return services.NewPublicationService(config, logger, redisClient)
}

func NewChannelService(logger *zap.Logger, redisClient redis.UniversalClient) services.ChannelService {
	return services.NewChannelService(logger, redisClient)
}

func NewPersistentChannelService(config *viper.Viper, logger *zap.Logger, publicationService services.PublicationService, channelService services.ChannelService) services.PersistentChannelService {
	return services.NewPersistentChannelService(config, logger, publicationService, channelService)
}
