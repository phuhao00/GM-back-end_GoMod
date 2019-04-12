package common
import (
	. "HA-back-end/models"
	. "HA-back-end/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

////
//func CheckToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		token := c.GetHeader("X-Token")
//		user := GetUserByToken(token)
//		if c.Request.URL.String() != "/api/xmw/pay/notify" && c.Request.URL.String() != "/api/_test/pay/start" && c.Request.URL.String() != "/api/_test2/pay/start" && c.Request.URL.String() != "/api/xmw/pay/start" && c.Request.URL.String() != "/api/admin/login" && c.Request.URL.String() != "/api/admin/serverList" && c.Request.URL.String() != "/api/track/server"  {
//			if len(token) == 0 || user.Token != token || time.Now().Unix() > user.TokenExpireTime.Unix() {
//				c.JSON(http.StatusUnauthorized, "token error")
//			}
//			if c.Keys == nil {
//				c.Keys = make(map[string]interface{})
//			}
//			c.Keys["name"] = user.Name
//			c.Keys["token"] = token
//		}
//		c.Next()
//	}
//}
//
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Request.Method == http.MethodPost && c.Request.URL.String() != "/api/track/server" {
			if c.Keys["name"] == nil {
				return
			}
			NewUserLog(&UserLog{
				Name:   c.Keys["name"].(string),
				Time:   time.Now(),
				Ip:     c.ClientIP(),
				Url:    c.Request.URL.String(),
				Status: int32(c.Writer.Status()),
			})
		}
	}
}