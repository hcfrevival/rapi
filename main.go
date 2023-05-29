package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rapi/config"
	"rapi/db"
	"rapi/routers"
)

func main() {
	// conf init
	conf := config.Prepare()
	gin.SetMode(conf.Gin.Mode)

	// db init
	mongoClient, err := db.CreateMongoClient(conf.Mongo.URI, conf.Mongo.DatabaseName)
	if err != nil {
		panic("failed to establish mongo connection: " + err.Error())
	}

	// middleware
	corsConf := cors.DefaultConfig()
	corsConf.AllowOrigins = conf.Gin.Origins
	corsConf.AllowCredentials = true
	corsConf.AddAllowHeaders("Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "Set-Cookie", "Access-Control-Allow-Origin")
	corsConf.AddAllowMethods("GET", "POST", "PUT", "DELETE", "OPTIONS")

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(corsConf))

	// route controller, passes mongo + db name to children
	rc := routers.RouteController{
		Mongo:        mongoClient,
		DatabaseName: conf.Mongo.DatabaseName,
	}

	// routes
	rc.ApplyLeaderboardRoutes(router)
	rc.ApplyFeedRoutes(router)
	rc.ApplyServerRoutes(router)

	// start gin
	err = router.Run(":" + conf.Gin.Port)
	if err != nil {
		panic("failed to start gin: " + err.Error())
	}
}
