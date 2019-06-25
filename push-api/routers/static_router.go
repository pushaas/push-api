package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type (
	StaticRouter interface {
		Router
	}

	staticRouter struct{
		staticsPath string
	}
)

func (r *staticRouter) SetupRoutes(router gin.IRouter) {
	router.Static("", r.staticsPath)
}

func NewStaticRouter(config *viper.Viper) Router {
	staticsPath := config.GetString("api.statics_path")

	return &staticRouter{
		staticsPath: staticsPath,
	}
}
