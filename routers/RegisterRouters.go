package routers

import (
	. "HA-back-end/routers/admin"
	"github.com/gin-gonic/gin"
)

var E *gin.Engine

func Init(){
	RegisterAdminRouters()
}
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