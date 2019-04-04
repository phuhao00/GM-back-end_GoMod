package admin

import (
	"HA-back-end/public"
	"HA-back-end/service"
	"github.com/gin-gonic/gin"
)
func CreateUser(c *gin.Context) {
	var request RegisterReq
	err := c.ShouldBindJSON(&request)
	if err == nil {
		err, user :=service.NewUser(request.Name, request.Password, request.Roles,request.Introduce); if err == nil {
			public.JsonSuccess(c, user)
			return
		}
	}
	public.JsonError(c, err)
}
