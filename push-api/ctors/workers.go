package ctors

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/services"
	"github.com/rafaeleyng/push-api/push-api/workers"
)

func NewPersistentChannelsWorker(lc fx.Lifecycle, config *viper.Viper, logger *zap.Logger, persistentChannelService services.PersistentChannelService) workers.PersistentChannelsWorker {
	return workers.NewPersistentChannelsWorker(lc, config, logger, persistentChannelService)
}
