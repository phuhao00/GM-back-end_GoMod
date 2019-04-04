package main

import (
	"HA-back-end/conf"
	"HA-back-end/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)
import "HA-back-end/DBMgr"

func main()  {
	conf.LoadProjectConf()//加载工程配置文件
	DBMgr.LoadMysqlConfig().OpenDB()//打开数据库连接
}


func Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	Run(routers.E, "127.0.0.1:7070")
}
//
func Run(router *gin.Engine, serverHost string) {

	server := &http.Server{
		Addr:    serverHost,
		Handler: router,
	}

	go func() {
		for {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	pid := fmt.Sprintf("%d", os.Getpid())
	_, openErr := os.OpenFile("pid", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openErr == nil {
		ioutil.WriteFile("pid", []byte(pid), 0)
	}
}