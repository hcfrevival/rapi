package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"rapi/config"
	"rapi/db"
	"rapi/model"
)

func (controller *Controller) GetKillFeed() gin.HandlerFunc {
	mqp := db.MongoQueryParams{
		MongoClient:    controller.Mongo,
		DatabaseName:   controller.DatabaseName,
		CollectionName: "stats_kills",
	}

	conf := config.Prepare()

	return func(ctx *gin.Context) {
		res, err := db.FindManyDocumentsByFilterWithOpts[model.Kill](
			mqp,
			bson.M{"map": conf.Factions.MapNumber},
			options.Find().SetLimit(10).SetSort(bson.D{{Key: "date", Value: -1}}),
		)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.AbortWithStatusJSON(http.StatusNotFound, CreateErrorResponse("no entries"))
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, CreateErrorResponse("failed to perform query"))
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (controller *Controller) GetDeathFeed() gin.HandlerFunc {
	mqp := db.MongoQueryParams{
		MongoClient:    controller.Mongo,
		DatabaseName:   controller.DatabaseName,
		CollectionName: "stats_deaths",
	}

	conf := config.Prepare()

	return func(ctx *gin.Context) {
		res, err := db.FindManyDocumentsByFilterWithOpts[model.Death](
			mqp,
			bson.M{"map": conf.Factions.MapNumber},
			options.Find().SetLimit(10).SetSort(bson.D{{Key: "date", Value: -1}}),
		)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.AbortWithStatusJSON(http.StatusNotFound, CreateErrorResponse("no entries"))
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, CreateErrorResponse("failed to perform query"))
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (controller *Controller) GetEventCaptureFeed() gin.HandlerFunc {
	mqp := db.MongoQueryParams{
		MongoClient:    controller.Mongo,
		DatabaseName:   controller.DatabaseName,
		CollectionName: "stats_events",
	}

	conf := config.Prepare()

	return func(ctx *gin.Context) {
		res, err := db.FindManyDocumentsByFilterWithOpts[model.EventCapture](
			mqp,
			bson.M{"map": conf.Factions.MapNumber},
			options.Find().SetLimit(10).SetSort(bson.D{{Key: "date", Value: -1}}),
		)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.AbortWithStatusJSON(http.StatusNotFound, CreateErrorResponse("no entries"))
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, CreateErrorResponse("failed to perform query"))
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
