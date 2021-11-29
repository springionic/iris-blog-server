// Package handler
// created by lilei at 2021/11/23
package handler

import (
	"github.com/kataras/iris/v12"
	"iris-blog-server/config"
)

type Response struct {
	Code uint     `json:"code"`
	Msg  string   `json:"msg"`
	Data iris.Map `json:"data,omitempty"`
}

type BaseHandler struct {
	Ctx iris.Context
}

// Success 成功不返回数据
func (c *BaseHandler) Success() Response {
	return Response{Code: 0, Msg: "success"}
}

// SuccessData 成功并返回数数据
func (c *BaseHandler) SuccessData(data iris.Map) Response {
	return Response{Code: 0, Msg: "success", Data: data}
}

// Error program execute error, `server error` and `client error` and `business error`
func (c *BaseHandler) Fail(err error) Response {
	var customError config.CustomError
	switch err.(type) {
	case config.CustomError:
		customError = err.(config.CustomError)
	default:
		customError = config.ServerIntervalError
	}
	c.Ctx.StatusCode(customError.HTTPCode())
	return Response{Code: customError.Code(), Msg: customError.Message()}
}
