package apiV1

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/pushaas/push-api/push-api/models"
	"github.com/pushaas/push-api/push-api/routers"
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
	pushStreamUrl := r.config.GetString("push_stream.url")
	parsedUrl, err := url.Parse(pushStreamUrl)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code: models.ErrorConfigGetParsePushStreamUrl,
			Message: "failed to parse push-stream URL",
		})
		return
	}

	config := map[string]interface{}{
		"pushStream": map[string]string {
			"url": pushStreamUrl,
			"hostname": parsedUrl.Hostname(),
			"port": parsedUrl.Port(),
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
