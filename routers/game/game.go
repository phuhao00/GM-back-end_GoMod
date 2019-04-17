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
			public.JsonSuccess(c, gin.H{"add game ": "success"})
			return
		}
		public.JsonError(c, error)

	}
}