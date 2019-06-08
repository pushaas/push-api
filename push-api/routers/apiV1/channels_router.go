package apiV1

import (
	"github.com/gin-gonic/gin"

	"github.com/rafaeleyng/push-api/push-api/routers"
)

type (
	ChannelsRouter interface {
		routers.Router
	}

	channelsRouter struct {
	}
)

func (r *channelsRouter) getChannels(c *gin.Context) {

}

func (r *channelsRouter) postChannel(c *gin.Context) {

}

func (r *channelsRouter) SetupRoutes(router gin.IRouter) {
	// service instance
	router.GET("/", r.getChannels)
	router.POST("/", r.postChannel)
}

func NewChannelsRouter() routers.Router {
	return &channelsRouter{}
}
