package apiV1

import (
	"net/http"
	"net/url"

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
	message.Content = url.PathEscape(message.Content)
	return &message, err
}

func (r *messagesRouter) postMessage(c *gin.Context) {
	message, err := messageFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Code: models.ErrorMessageCreateInvalidBody,
			Message: "invalid request body",
		})
		return
	}

	result := r.publicationService.PublishMessage(message)

	if result == services.PublishingFailure {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code: models.ErrorMessageCreateFailed,
			Message: "could not send message",
		})
		return
	}

	if result == services.PublishingInvalid {
		c.JSON(http.StatusBadRequest, models.Error{
			Code: models.ErrorMessageCreateInvalidMessageFormat,
			Message: "invalid message format",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (r *messagesRouter) SetupRoutes(router gin.IRouter) {
	router.POST("", r.postMessage)
}

func NewMessagesRouter(publicationService services.PublicationService) MessagesRouter {
	return &messagesRouter{
		publicationService: publicationService,
	}
}
