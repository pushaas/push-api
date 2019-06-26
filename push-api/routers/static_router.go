package routers

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type (
	StaticRouter interface {
		Router
		SetupClientSideRoutesSupport(engine *gin.Engine)
	}

	staticRouter struct{
		indexPath string
		staticsPath string
	}
)

func (r *staticRouter) SetupClientSideRoutesSupport(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		c.File(r.indexPath)
	})
}

func (r *staticRouter) SetupRoutes(router gin.IRouter) {
	router.Static("", r.staticsPath)
}

func NewStaticRouter(config *viper.Viper) StaticRouter {
	staticsPath := config.GetString("api.statics_path")
	indexPath := filepath.Join(staticsPath, "index.html")

	return &staticRouter{
		indexPath: indexPath,
		staticsPath: staticsPath,
	}
}
