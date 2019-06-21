package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rafaeleyng/push-api/push-api/models"
	"github.com/rafaeleyng/push-api/push-api/routers"
	"github.com/rafaeleyng/push-api/push-api/services"
)

type (
	MessagesRouter interface {
		routers.Router
	}

	messagesRouter struct {
		publicationService services.PublicationService
	}
)

func messageFromContext(c *gin.Context) (*models.Message, error) {
	var message models.Message
	err := c.BindJSON(&message)
	if err != nil {
		return nil, err
	}
	return &message, err
}

func (r *messagesRouter) postMessage(c *gin.Context) {
	message, err := messageFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			// TODO add remaining fields
			Message: "invalid request body",
		})
		return
	}

	result := r.publicationService.PublishMessage(message)

	if result == services.MessagePublishingFailure {
		c.JSON(http.StatusInternalServerError, models.Error{
			// TODO add remaining fields
			Message: "could not send message",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (r *messagesRouter) SetupRoutes(router gin.IRouter) {
	router.POST("", r.postMessage)
}

func NewMessagesRouter(publicationService services.PublicationService) routers.Router {
	return &messagesRouter{
		publicationService: publicationService,
	}
}
