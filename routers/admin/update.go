package admin

import (
	"HA-back-end/public"
	"HA-back-end/service"
	"github.com/gin-gonic/gin"
)
//
func Update(c *gin.Context)  {
	var request UpdateReq
	err := c.ShouldBindJSON(&request)
	if err == nil {
		err, user :=service.UpdateUser(request.Name,request.ColumnName,request.ColumnVal); if err == nil {
			public.JsonSuccess(c, user)
			return
		}
	}
	public.JsonError(c, err)
}
