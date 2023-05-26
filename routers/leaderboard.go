package routers

import (
	"github.com/gin-gonic/gin"
	"rapi/controllers"
	"rapi/model"
)

// ApplyLeaderboardRoutes applies all leaderboard endpoints to the
// provided gin router instance
func (rc *RouteController) ApplyLeaderboardRoutes(router *gin.Engine) {
	c := controllers.Controller{
		Mongo:        rc.Mongo,
		DatabaseName: rc.DatabaseName,
	}

	v1 := router.Group("/v1/leaderboard")
	{
		v1.GET("/kills", c.GetLeaderboardQuery(model.KILL))
		v1.GET("/deaths", c.GetLeaderboardQuery(model.DEATH))
		v1.GET("/events", c.GetLeaderboardQuery(model.EVENT_CAPTURES))
		v1.GET("/longshot", c.GetLeaderboardQuery(model.LONGSHOT))
		v1.GET("/exp", c.GetLeaderboardQuery(model.EXP_EARNED))
		v1.GET("/playtime", c.GetLeaderboardQuery(model.PLAYTIME))
	}
}
