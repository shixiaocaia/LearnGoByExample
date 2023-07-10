package router

import (
	"github.com/gin-gonic/gin"
	"go_example/web_app/logger"
	"go_example/web_app/settings"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r

}
