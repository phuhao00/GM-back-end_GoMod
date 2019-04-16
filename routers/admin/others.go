package admin

import (
	. "HA-back-end/models"
	"HA-back-end/public"
	"HA-back-end/service"
	"github.com/gin-gonic/gin"
)

//
func GetUserPlayedGames(c *gin.Context) {
	var request GetUserHavePlayGamesReq
	err := c.ShouldBindJSON(&request)
	if err == nil {
		err, gameIds:= service.GetGamesInfo(request.Name)
		if err!=nil {

		}
		var games []*Game =make([]*Game,0)
		games=service.GetAllGameInfos(gameIds)
		public.JsonSuccess(c, gin.H{"gamesInfo": games})
	}
}
