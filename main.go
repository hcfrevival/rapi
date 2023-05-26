package main

import (
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
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// route controller, passes mongo + db name to children
	rc := routers.RouteController{
		Mongo:        mongoClient,
		DatabaseName: conf.Mongo.DatabaseName,
	}

	// routes
	rc.ApplyLeaderboardRoutes(router)
	rc.ApplyFeedRoutes(router)

	// start gin
	err = router.Run(":" + conf.Gin.Port)
	if err != nil {
		panic("failed to start gin: " + err.Error())
	}
}
