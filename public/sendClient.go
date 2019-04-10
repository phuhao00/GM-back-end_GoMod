package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message":  "ok",
		"data": data,
	})
}

func JsonError(c *gin.Context, msg error) {
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"message":  msg,
	})
}

func JsonFail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"message":  msg,
	})
}