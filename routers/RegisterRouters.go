package routers

import (
	. "HA-back-end/routers/Test"
	. "HA-back-end/routers/admin"
	. "HA-back-end/routers/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

var E *gin.Engine

func Init(){

	E=gin.New()
	E.Use(gin.Logger())
	E.Use(gin.Recovery())
	E.Use(CheckToken())
	E.Use(Logger())
	E.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})
	RegisterAdminRouters()
	RegisterTestRouters()
}
//
func RegisterTestRouters()  {
	action:= E.Group("/")
	{
		action:=action.Group("/test")
		{
			action.POST("/1",Test)
		}
	}
}
//
func RegisterAdminRouters()  {
	action:= E.Group("/")
	{
		action:=action.Group("/admin")
		{
			action.POST("login",Login)
			action.POST("register", CreateUser)
		}
	}
}