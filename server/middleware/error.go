// Package middleware
// created by lilei at 2021/11/24
package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"iris-blog-server/util"
	"runtime"
)

var errorLogger = util.NewLoggerWriter(logrus.InfoLevel, "error", true)

func ErrorHandler(c iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			errMsg := fmt.Sprintf("\nerror: %s", err)
			for i := 1; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				errMsg += fmt.Sprintf("\n%s:%d", file, line)
			}
			errorLogger.Error(errMsg)
			c.StopWithJSON(iris.StatusInternalServerError, iris.Map{
				"code": iris.StatusInternalServerError,
				"msg":  "服务器内部错误",
			})
		}
	}()
}
