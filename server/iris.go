// Package server
//created by lilei at 2021/10/7
package server

import (
	"context"
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"iris-blog-server/comp"
	"iris-blog-server/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	log.Println("start server...")
	app := iris.New()
	app.Use(middleware.LoggerHandler)
	regisRouter(app)
	err := app.Run(iris.Addr(":8080"), iris.WithConfiguration(
		iris.Configuration{
			DisableStartupLog:                 false,
			DisableInterruptHandler:           false,
			DisablePathCorrection:             false,
			EnablePathEscape:                  false,
			FireMethodNotAllowed:              false,
			DisableBodyConsumptionOnUnmarshal: false,
			DisableAutoFireStatusCode:         false,
			EnableOptimizations:               true,
			Charset:                           "UTF-8",
			TimeFormat:                        "2006-01-02 15:04:05",
		},
	))

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	// 反之，则执行后续代码
	<-quit
	signal.Reset(syscall.SIGINT, syscall.SIGTERM)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用 app.Shutdown() 完成优雅停止
	// 调用时传递了一个上下文对象，对象中定义了超时时间
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// 终止所有日志记录器
	for i := 0; i < len(comp.LoggerWriterList); i++ {
		comp.LoggerWriterList[i].Stop()
	}
	log.Println("Server exited.")
}
