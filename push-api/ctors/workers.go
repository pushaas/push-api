package ctors

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/pushaas/push-api/push-api/services"
	"github.com/pushaas/push-api/push-api/workers"
)

func NewPersistentChannelsWorker(lc fx.Lifecycle, config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient, persistentChannelService services.PersistentChannelService) workers.PersistentChannelsWorker {
	return workers.NewPersistentChannelsWorker(lc, config, logger, redisClient, persistentChannelService)
}
