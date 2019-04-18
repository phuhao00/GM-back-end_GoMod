package game

import (
	"HA-back-end/models"
	"HA-back-end/public"
	"HA-back-end/service"
	"github.com/gin-gonic/gin"
)
import ."HA-back-end/routers/admin"

//添加游戏
func AddGame(c *gin.Context)  {
var request	GameAddParamReq
	err := c.ShouldBindJSON(&request)
	if err == nil {
		game:=&models.Game{
			Name: request.Name,
			Price:request.Price,
			URL:request.Url,
			SupplierID:request.SupplierId,
		}
		error := service.InsertGame(game)
		if error!=nil {
			public.JsonError(c, error)
			return
		}
		public.JsonSuccess(c, gin.H{"add game ": "success"})
	}
}
//
func UpdateGame(c *gin.Context)  {
	var request GameUpdateReqParam
	err := c.ShouldBindJSON(&request)
	if err == nil {
		game:=&models.Game{
			ID:               request.ID,
			Name:             request.Name,
			Price:            request.Price,
			CommentID:        request.CommentID,
			DownloadQuantity: request.DownloadQuantity,
			Score:            request.Score,
			URL:              request.URL,
			SupplierID:       request.SupplierID,
		}
		error := service.UpdateGame(game)
		if error!=nil {
			public.JsonError(c, error)
			return
		}
		public.JsonSuccess(c, gin.H{"UpdateGame ": "success"})
	}
}
//
func DeleteGame(c *gin.Context)  {
	var  request  GameDeleteReqParam
	err := c.ShouldBindJSON(&request)
	if err == nil {
		error:=service.DeleteGame(request.ID)
		if error!=nil {
			public.JsonError(c, error)
			return
		}
		public.JsonSuccess(c, gin.H{"DeleteGame ": "success"})
	}
}