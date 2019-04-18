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

//
func InsertGame(game *Game)(err error)  {
	 dbc:=DBMgr.MySql.Save(&game)
	if dbc.Error != nil {
		return dbc.Error
	}
	 return nil
}

func UpdateGame(game *Game) (err error) {
	g:=&Game{}
	dbc:=DBMgr.MySql.Model(&g).Where("id=?",game.ID).Updates(map[string]interface{}{
		//"name": game.Name,
		"price": game.Price,
		"commentId": game.CommentID,
		"downloadQuantity":game.DownloadQuantity,
		"score":game.Score,
		"url":game.URL,
		//"supplierId":game.SupplierID,
	})
	if dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func DeleteGame(GameId int )(err error)  {
	g:=&Game{}
	dbc:=DBMgr.MySql.Where("id=?", GameId).Delete(g)
	if dbc.Error != nil {
		return dbc.Error
	}
	return nil
}