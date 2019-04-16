package service

import
(
	."HA-back-end/models"
)
//
func GetAllGameInfos(gameIds []int64)[]*Game  {
	for _,ele:=range gameIds{
		GetGameInfo(ele)
	}
	return nil
}

func GetGameInfo(gameId int64) *Game  {


	return nil
}