package main

import (
	"HA-back-end/DBMgr"
	."HA-back-end/ServerMgr"
	."HA-back-end/conf"
	."HA-back-end/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"runtime"
)
func main()  {
	// 创建记录日志的文件
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	LoadProjectConf()//加载工程配置文件
	DBMgr.LoadMysqlConfig().OpenDB()//打开数据库连接
	Start()
}
//
func Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	routers:=InitRouter()
	RunBaseWebModule(routers, Mode["Dev"].Addr+":"+Mode["Dev"].Port)
	ServerManager.RunClient(Mode["Dev"].Addr)
}