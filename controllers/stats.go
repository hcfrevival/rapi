package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"rapi/config"
	"rapi/db"
	"rapi/model"
	"strconv"
	"strings"
)

// GetLeaderboardQuery performs a leaderboard query for the provided statistic type
// and returns it in a sanitized JSON response
func (controller *Controller) GetLeaderboardQuery(statType model.EStatisticType) gin.HandlerFunc {
	type SanitizedStats struct {
		UUID     string `json:"uuid"`
		Username string `json:"username"`
		Value    uint64 `json:"value"`
	}

	mqp := db.MongoQueryParams{
		MongoClient:    controller.Mongo,
		DatabaseName:   controller.DatabaseName,
		CollectionName: "stats_players",
	}

	accMqp := db.MongoQueryParams{
		MongoClient:    controller.Mongo,
		DatabaseName:   controller.DatabaseName,
		CollectionName: "accounts",
	}

	conf := config.Prepare()

	return func(ctx *gin.Context) {
		var page = 0
		var err error

		pageStr, pagePresent := ctx.GetQuery("page")
		if pagePresent {
			page, err = strconv.Atoi(pageStr)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, CreateErrorResponse("invalid page number"))
				return
			}
		}

		res, err := db.FindManyDocumentsByFilterWithOpts[model.StatHolder](
			mqp,
			bson.M{"map": conf.Factions.MapNumber},
			options.Find().SetLimit(10).SetSkip(int64(page*10)).SetSort(bson.D{{Key: strings.ToUpper(string(statType)), Value: -1}}),
		)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, CreateErrorResponse("failed to perform query: "+err.Error()))
			return
		}

		var sanitized []SanitizedStats
		for _, doc := range res {
			aresAccount, err := db.FindDocumentByKeyValue[string, model.Account](accMqp, "uuid", doc.UUID)
			if err != nil {
				continue
			}

			var v uint64 = 0
			switch {
			case statType == model.KILL:
				v = doc.Kills
				break
			case statType == model.DEATH:
				v = doc.Deaths
				break
			case statType == model.EVENT_CAPTURES:
				v = doc.EventCaptures
				break
			case statType == model.EXP_EARNED:
				v = doc.ExpEarned
				break
			case statType == model.LONGSHOT:
				v = doc.Longshot
				break
			case statType == model.PLAYTIME:
				v = doc.Playtime
				break
			}

			s := SanitizedStats{
				UUID:     doc.UUID,
				Username: aresAccount.Username,
				Value:    v,
			}

			sanitized = append(sanitized, s)
		}

		ctx.JSON(http.StatusOK, sanitized)
	}
}
