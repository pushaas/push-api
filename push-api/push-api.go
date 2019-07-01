package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/ctors"
	"github.com/rafaeleyng/push-api/push-api/workers"
)

func runApp(logger *zap.Logger, router *gin.Engine, config *viper.Viper, persistentChannelsWorker workers.PersistentChannelsWorker) error {
	log := logger.Named("runApp")

	persistentChannelsWorker.DispatchWorker()

	err := router.Run(fmt.Sprintf(":%s", config.GetString("server.port")))
	if err != nil {
		log.Error("error on running server", zap.Error(err))
		return err
	}

	return nil
}

func Run() {
	app := fx.New(
		fx.Provide(
			ctors.NewViper,
			ctors.NewLogger,
			ctors.NewRedisClient,

			// routers
			ctors.NewGinRouter,
			ctors.NewStaticRouter,
			ctors.NewApiRootRouter,
			ctors.NewAuthRouter,
			ctors.NewChannelsRouter,
			ctors.NewMessagesRouter,
			ctors.NewStatsRouter,

			// services
			ctors.NewChannelService,
			ctors.NewPersistentChannelService,
			ctors.NewPublicationService,
			ctors.NewStatsService,

			// workers
			ctors.NewPersistentChannelsWorker,
		),
		fx.Invoke(runApp),
	)

	app.Run()
}
