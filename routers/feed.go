package routers

import (
	"github.com/gin-gonic/gin"
	"rapi/controllers"
)

func (rc *RouteController) ApplyFeedRoutes(router *gin.Engine) {
	c := controllers.Controller{
		Mongo:          rc.Mongo,
		DatabaseName:   rc.DatabaseName,
		CollectionName: "",
	}

	v1 := router.Group("/v1/feed")
	{
		v1.GET("/kills", c.GetKillFeed())
		v1.GET("/deaths", c.GetDeathFeed())
		v1.GET("/events", c.GetEventCaptureFeed())
	}
}
