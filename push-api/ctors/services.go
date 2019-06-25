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

func NewChannelService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) services.ChannelService {
	return services.NewChannelService(config, logger, redisClient)
}

func NewPersistentChannelService(logger *zap.Logger, publicationService services.PublicationService, channelService services.ChannelService) services.PersistentChannelService {
	return services.NewPersistentChannelService(logger, publicationService, channelService)
}

func NewStatsService(config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient) services.StatsService {
	return services.NewStatsService(config, logger, redisClient)
}
