package service

import
(
	"HA-back-end/DBMgr"
	."HA-back-end/models"
	"strconv"
)
//
func GetAllGameInfos(gameIdArr []string)(gamesInfo []*Game)  {
	gamesInfo=make([]*Game,0)
	for _,ele:=range gameIdArr{
		s, err :=strconv.ParseInt(ele,10,64)
		if err==nil {
			gamesInfo=append(gamesInfo,GetGameInfo(s))
		}
	}
	return
}

func GetGameInfo(gameId int64) *Game  {
	game:=&Game{}
	dbc:=DBMgr.MySql.Where("id=?",gameId).First(&game)
	if dbc.Error != nil {
		return nil
	}
	return game
}