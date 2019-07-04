package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rafaeleyng/push-api/push-api/models"
	"github.com/rafaeleyng/push-api/push-api/routers"
	"github.com/rafaeleyng/push-api/push-api/services"
)

type (
	ChannelsRouter interface {
		routers.Router
	}

	channelsRouter struct {
		channelService services.ChannelService
	}
)

func channelFromContext(c *gin.Context) (*models.Channel, error) {
	var channel models.Channel
	err := c.BindJSON(&channel)
	if err != nil {
		return nil, err
	}
	return &channel, err
}

func (r *channelsRouter) postChannel(c *gin.Context) {
	channel, err := channelFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			// TODO add remaining fields
			Message: "invalid request body",
		})
		return
	}

	result := r.channelService.Create(channel)

	if result == services.ChannelCreationAlreadyExist {
		c.JSON(http.StatusBadRequest, models.Error{
			// TODO add remaining fields
			Message: "a channel with this id already exists",
		})
		return
	}

	if result == services.ChannelCreationFailure {
		c.JSON(http.StatusInternalServerError, models.Error{
			// TODO add remaining fields
			Message: "failed to create channel",
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (r *channelsRouter) getChannel(c *gin.Context) {
	id := c.Param("id")
	channel, result := r.channelService.Get(id)

	if result == services.ChannelRetrievalNotFound {
		c.JSON(http.StatusNotFound, models.Error{
			// TODO add remaining fields
			Message: "channel not found",
		})
		return
	}

	if result == services.ChannelRetrievalFailure {
		c.JSON(http.StatusInternalServerError, models.Error{
			// TODO add remaining fields
			Message: "failed to retrieve channel",
		})
		return
	}

	c.JSON(http.StatusOK, channel)
}

func (r *channelsRouter) deleteChannel(c *gin.Context) {
	id := c.Param("id")
	result := r.channelService.Delete(id)

	if result == services.ChannelDeletionNotFound {
		c.JSON(http.StatusNotFound, models.Error{
			// TODO add remaining fields
			Message: "channel not found",
		})
		return
	}

	if result == services.ChannelDeletionFailure {
		c.JSON(http.StatusInternalServerError, models.Error{
			// TODO add remaining fields
			Message: "failed to delete channel",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (r *channelsRouter) getChannels(c *gin.Context) {
	channels, result := r.channelService.GetAll()

	if result == services.ChannelRetrievalFailure {
		c.JSON(http.StatusInternalServerError, models.Error{
			// TODO add remaining fields
			Message: "failed to retrieve channels",
		})
	}

	c.JSON(http.StatusOK, channels)
}

func (r *channelsRouter) SetupRoutes(router gin.IRouter) {
	router.POST("", r.postChannel)
	router.GET("/:id", r.getChannel)
	router.DELETE("/:id", r.deleteChannel)
	router.GET("", r.getChannels)
}

func NewChannelsRouter(channelService services.ChannelService) ChannelsRouter {
	return &channelsRouter{
		channelService: channelService,
	}
}
