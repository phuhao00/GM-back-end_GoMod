package admin

import (
	"HA-back-end/public"
	"HA-back-end/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	var request LoginReq
	err := c.ShouldBindJSON(&request)
		if err == nil {
		err, token := service.Login(request.UserName, request.Password, c.ClientIP())
		if err == nil {
			if c.Keys == nil {
				c.Keys = make(map[string]interface{})
			}
			c.Keys["name"] = request.UserName
			c.Keys["token"] = token
			public.JsonSuccess(c, gin.H{"token": token})
			return
		}
	}

	public.JsonError(c, err)
}