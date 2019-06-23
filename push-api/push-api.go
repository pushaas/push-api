package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/rafaeleyng/push-api/push-api/ctors"
	"github.com/rafaeleyng/push-api/push-api/workers"
)

func runApp(router *gin.Engine, config *viper.Viper, persistentChannelsWorker workers.PersistentChannelsWorker) error {
	persistentChannelsWorker.DispatchWorker()

	err := router.Run(fmt.Sprintf(":%s", config.GetString("server.port")))
	return err
}

func Run() {
	app := fx.New(
		fx.Provide(
			ctors.NewViper,
			ctors.NewLogger,
			ctors.NewRedisClient,
			ctors.NewMachineryServer,

			// routers
			ctors.NewRouter,
			ctors.NewRootRouter,
			ctors.NewStaticRouter,
			ctors.NewApiRootRouter,
			ctors.NewChannelsRouter,
			ctors.NewMessagesRouter,

			// services
			ctors.NewChannelService,
			ctors.NewPersistentChannelService,
			ctors.NewPublicationService,

			// workers
			ctors.NewPersistentChannelsWorker,
		),
		fx.Invoke(runApp),
	)

	app.Run()
}
