package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/rafaeleyng/push-api/push-api/routers"
)

type (
	ConfigRouter interface {
		routers.Router
	}

	configRouter struct {
		config *viper.Viper
	}
)

func (r *configRouter) getConfig(c *gin.Context) {
	config := map[string]interface{}{
		"pushStream": map[string]string {
			"url": r.config.GetString("push_stream.url"),
		},
	}

	c.JSON(http.StatusOK, config)
}

func (r *configRouter) SetupRoutes(router gin.IRouter) {
	router.GET("", r.getConfig)
}

func NewConfigRouter(config *viper.Viper) ConfigRouter {
	return &configRouter{
		config: config,
	}
}
