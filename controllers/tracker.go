package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"rapi/db"
	"rapi/model"
)

// GetTracker reads a provided 'id' param
// and queries the database for tracker data
func (controller *Controller) GetTracker() gin.HandlerFunc {
	mqp := db.MongoQueryParams{
		MongoClient:    controller.Mongo,
		DatabaseName:   controller.DatabaseName,
		CollectionName: controller.CollectionName,
	}
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		res, err := db.FindDocumentById[model.EventTracker](mqp, id)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.AbortWithStatus(http.StatusNotFound)
				return
			}

			fmt.Println("failed to query event tracker", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, CreateErrorResponse("failed to query documents"))
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
