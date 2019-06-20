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

func NewRouter(
	config *viper.Viper,
	logger *zap.Logger,
	rootRouter routers.RootRouter,
	staticRouter routers.StaticRouter,
	apiRootRouter routers.ApiRootRouter,
	v1ChannelsRouter apiV1.ChannelsRouter,
	v1MessagesRouter apiV1.MessagesRouter,
) *gin.Engine {
	envConfig := config.Get("env")
	if envConfig == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	g(r, "/", func(r gin.IRouter) {
		rootRouter.SetupRoutes(r)
	})

	g(r, "/static", func(r gin.IRouter) {
		staticRouter.SetupRoutes(r)
	})

	g(r, "/api", func(r gin.IRouter) {
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
		})
	})

	return r
}

func NewRootRouter() routers.RootRouter {
	return routers.NewRootRouter()
}

func NewStaticRouter() routers.StaticRouter {
	return routers.NewStaticRouter()
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
