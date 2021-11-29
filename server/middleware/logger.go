// Package middleware
//created by lilei at 2021/10/31
package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"iris-blog-server/util"
	"time"
)

var logger = util.NewLoggerWriter(log.InfoLevel, "access.log", true)

type AccessLog struct {
	StartTime   time.Time
	EndTime     time.Time
	StatusCode  int
	LatencyTime time.Duration
	ClientIP    string
	Svc         string
	Uid         string
	ReqMethod   string
	ReqUri      string
	ReqBody     []byte
	RspBody     []byte
}

func (a AccessLog) Format() string {
	return fmt.Sprintf(
		"%25s:::%25s::: %3d ::: %13v ::: %15s ::: %16s ::: %8s ::: %4s ::: %s ::: %s ::: %s :::",
		a.StartTime.Format("2006-01-02T15:04:05Z07:00"),
		a.EndTime.Format("2006-01-02T15:04:05Z07:00"),
		a.StatusCode, a.LatencyTime, a.ClientIP,
		a.Svc, a.Uid, a.ReqMethod, a.ReqUri, a.ReqBody, a.RspBody,
	)
}

func LoggerHandler(c iris.Context) {
	// 开始时间
	startTime := time.Now()

	// 请求Body
	reqBody, _ := c.GetBody()

	// 处理请求
	c.Next()

	// 请求路由
	reqUri := c.Path()

	// 请求方式
	reqMethod := c.Method()

	// 请求IP
	clientIP := c.RemoteAddr()

	// 状态码
	statusCode := c.GetStatusCode()

	// 响应Body
	rspBody := c.Recorder().Body()

	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)

	accessLog := AccessLog{
		StartTime:   startTime,
		EndTime:     endTime,
		StatusCode:  statusCode,
		LatencyTime: latencyTime,
		ClientIP:    clientIP,
		ReqMethod:   reqMethod,
		ReqUri:      reqUri,
		ReqBody:     reqBody,
		RspBody:     rspBody,
	}

	logger.Info(accessLog.Format())
}
