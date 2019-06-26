package ctors

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/routers/apiV1"
	"github.com/rafaeleyng/push-api/push-api/services"

	"github.com/rafaeleyng/push-api/push-api/routers"
)

func g(router gin.IRouter, path string, groupFn func(r gin.IRouter)) {
	groupFn(router.Group(path))
}

func getNoAuthMiddleware(config *viper.Viper, logger *zap.Logger) gin.HandlerFunc {
	logger.Debug("configuring no auth middleware")

	return func(c *gin.Context) {}
}

func getBasicAuthMiddleware(config *viper.Viper, logger *zap.Logger) gin.HandlerFunc {
	user := config.GetString("api.basic_auth_user")
	password := config.GetString("api.basic_auth_password")

	logger.Debug("configuring basic auth middleware", zap.String("user", user), zap.String("password", password))

	return gin.BasicAuth(gin.Accounts{
		user: password,
	})
}

func getAuthMiddleware(config *viper.Viper, logger *zap.Logger) gin.HandlerFunc {
	if enableAuth := config.GetBool("api.enable_auth"); enableAuth {
		return getBasicAuthMiddleware(config, logger)
	}

	return getNoAuthMiddleware(config, logger)
}

func NewGinRouter(
	config *viper.Viper,
	logger *zap.Logger,

	// static
	staticRouter routers.StaticRouter,

	// api
	apiRootRouter routers.ApiRootRouter,

	// api v1
	v1ChannelsRouter apiV1.ChannelsRouter,
	v1MessagesRouter apiV1.MessagesRouter,
	v1StatsRouter apiV1.StatsRouter,
) *gin.Engine {
	envConfig := config.Get("env")
	if envConfig == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	baseRouter := gin.Default()

	g(baseRouter, "/api", func(r gin.IRouter) {
		r.Use(getAuthMiddleware(config, logger))

		g(r, "/", func(r gin.IRouter) {
			apiRootRouter.SetupRoutes(r)
		})

		g(r, "/v1", func(r gin.IRouter) {
			g(r, "/channels", func(r gin.IRouter) {
				v1ChannelsRouter.SetupRoutes(r)
			})

			g(r, "/messages", func(r gin.IRouter) {
				v1MessagesRouter.SetupRoutes(r)
			})

			g(r, "/stats", func(r gin.IRouter) {
				v1StatsRouter.SetupRoutes(r)
			})
		})
	})

	g(baseRouter, "/admin", func(r gin.IRouter) {
		staticRouter.SetupRoutes(r)
		staticRouter.SetupClientSideRoutesSupport(baseRouter)
	})

	return baseRouter
}

func NewStaticRouter(config *viper.Viper) routers.StaticRouter {
	return routers.NewStaticRouter(config)
}

func NewApiRootRouter() routers.ApiRootRouter {
	return routers.NewApiRootRouter()
}

func NewChannelsRouter(channelService services.ChannelService) apiV1.ChannelsRouter {
	return apiV1.NewChannelsRouter(channelService)
}

func NewMessagesRouter(publicationService services.PublicationService) apiV1.MessagesRouter {
	return apiV1.NewMessagesRouter(publicationService)
}

func NewStatsRouter(statsService services.StatsService) apiV1.StatsRouter {
	return apiV1.NewStatsRouter(statsService)
}
