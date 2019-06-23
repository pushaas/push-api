package ctors

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/services"
)

func NewPublicationService(config *viper.Viper, logger *zap.Logger, machineryServer *machinery.Server) services.PublicationService {
	return services.NewPublicationService(config, logger, machineryServer)
}

func NewChannelService(logger *zap.Logger, redisClient redis.UniversalClient) services.ChannelService {
	return services.NewChannelService(logger, redisClient)
}

func NewPersistentChannelService(logger *zap.Logger, publicationService services.PublicationService, channelService services.ChannelService) services.PersistentChannelService {
	return services.NewPersistentChannelService(logger, publicationService, channelService)
}
