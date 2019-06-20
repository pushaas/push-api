package apiV1

import (
	"github.com/gin-gonic/gin"

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

func (r *messagesRouter) postMessage(c *gin.Context) {
	err := r.publicationService.Publish()
	if err != nil {
		// TODO
	}
}

func (r *messagesRouter) SetupRoutes(router gin.IRouter) {
	router.POST("/", r.postMessage)
}

func NewMessagesRouter(publicationService services.PublicationService) routers.Router {
	return &messagesRouter{
		publicationService: publicationService,
	}
}
