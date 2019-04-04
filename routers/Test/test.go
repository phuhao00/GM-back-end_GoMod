package Test

import (
	"HA-back-end/public"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context)  {
	public.JsonSuccess(c, gin.H{"token":" HelloWorld"})
}