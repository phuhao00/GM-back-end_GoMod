package routers

import (
	. "HA-back-end/routers/Test"
	. "HA-back-end/routers/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)


func InitRouter() *gin.Engine{
	var E *gin.Engine
	E=gin.New()
	//E.Use(gin.Logger())
	//E.Use(gin.Recovery())
	//E.Use(CheckToken())
	//E.Use(Logger())
	E.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})
	RegisterAdminRouters(E)
	RegisterTestRouters(E)
	return E
}
//
func RegisterTestRouters(E *gin.Engine)  {
	action:= E.Group("/user")
	{
		action:=action.Group("/test")
		{
			action.POST("/hh",Test)
		}
	}
}
//
func RegisterAdminRouters(E *gin.Engine)  {
	action:= E.Group("/user")
	{
		action:=action.Group("/admin")
		{
			action.POST("login",Login)
			action.POST("register", CreateUser)
		}
	}
}