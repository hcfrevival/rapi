package routers

import (
	"github.com/gin-gonic/gin"
	"rapi/controllers"
)

func (rc *RouteController) ApplyTrackerRoutes(router *gin.Engine) {
	c := controllers.Controller{
		Mongo:          rc.Mongo,
		DatabaseName:   rc.DatabaseName,
		CollectionName: "event_trackers",
	}

	v1 := router.Group("/v1/tracker")
	{
		v1.GET("/:id", c.GetTracker())
	}
}
