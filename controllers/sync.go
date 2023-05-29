package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"rapi/db"
	"rapi/model"
)

func (controller *Controller) GetServers() gin.HandlerFunc {
	mqp := db.MongoQueryParams{
		MongoClient:    controller.Mongo,
		DatabaseName:   controller.DatabaseName,
		CollectionName: controller.CollectionName,
	}

	return func(ctx *gin.Context) {
		res, err := db.FindManyDocumentsByFilter[model.SyncServer](mqp, bson.M{})
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.AbortWithStatusJSON(http.StatusNotFound, CreateErrorResponse("no servers found"))
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, CreateErrorResponse("failed to perform server query: "+err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
