package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pushaas/push-api/push-api/models"
	"github.com/pushaas/push-api/push-api/routers"
	"github.com/pushaas/push-api/push-api/services"
)

type (
	StatsRouter interface {
		routers.Router
	}

	statsRouter struct {
		statsService services.StatsService
	}
)

func (r *statsRouter) handleStatsRequest(c *gin.Context, stats interface{}, result services.StatsRetrievalResult) {
	if result == services.StatsRetrievalNotFound {
		c.JSON(http.StatusNotFound, models.Error{
			Code: models.ErrorStatsGetNotFound,
			Message: "stats data not found",
		})
		return
	}

	if result == services.StatsRetrievalFailure {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code: models.ErrorStatsGetFailed,
			Message: "failed to retrieve stats",
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (r *statsRouter) getGlobalStats(c *gin.Context) {
	stats, result := r.statsService.GetGlobalStats()
	r.handleStatsRequest(c, stats, result)
}

func (r *statsRouter) getChannelStats(c *gin.Context) {
	stats, result := r.statsService.GetChannelStats(c.Param("id"))
	r.handleStatsRequest(c, stats, result)
}

func (r *statsRouter) SetupRoutes(router gin.IRouter) {
	router.GET("/global", r.getGlobalStats)
	router.GET("/channels/:id", r.getChannelStats)
}

func NewStatsRouter(statsService services.StatsService) StatsRouter {
	return &statsRouter{
		statsService: statsService,
	}
}
