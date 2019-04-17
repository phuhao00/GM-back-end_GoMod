package admin

import (
	. "HA-back-end/models"
	"HA-back-end/public"
	"HA-back-end/service"
	"github.com/gin-gonic/gin"
	"strings"
)

//
func GetUserPlayedGames(c *gin.Context) {
	var request GetUserHavePlayGamesReq
	err := c.ShouldBindJSON(&request)
	if err == nil {
		err, gameIds:= service.GetGamesInfo(request.Name)
		if err!=nil {

		}
		var gameIdArr  []string
		if gameIds!=""{
			if strings.Contains(gameIds,",")  {
				gameIdArr=strings.Split(gameIds,",")
			}else {
				gameIdArr=append(gameIdArr,gameIds)
			}
		}
		var games []*Game =make([]*Game,0)
		games=service.GetAllGameInfos(gameIdArr)
		public.JsonSuccess(c, gin.H{"gamesInfo": games})
	}
}
