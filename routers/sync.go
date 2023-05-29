package routers

import (
	"github.com/gin-gonic/gin"
	"rapi/controllers"
)

func (rc *RouteController) ApplyServerRoutes(router *gin.Engine) {
	c := controllers.Controller{
		Mongo:          rc.Mongo,
		DatabaseName:   rc.DatabaseName,
		CollectionName: "sync",
	}

	v1 := router.Group("/v1/servers")
	{
		v1.GET("/", c.GetServers())
	}
}
