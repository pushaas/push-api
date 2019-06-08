package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/rafaeleyng/push-api/push-api/ctors"
)

func runApp(router *gin.Engine, config *viper.Viper) error {
	port := config.GetString("server.port")
	err := router.Run(fmt.Sprintf(":%s", port))
	return err
}

func Run() {
	app := fx.New(
		fx.Provide(
			ctors.NewViper,
			ctors.NewLogger,
			ctors.NewMongodb,

			// routers
			ctors.NewRouter,
			ctors.NewRootRouter,
			ctors.NewStaticRouter,
			ctors.NewApiRootRouter,
			ctors.NewChannelsRouter,

			// services
		),
		fx.Invoke(runApp),
	)

	app.Run()
}
