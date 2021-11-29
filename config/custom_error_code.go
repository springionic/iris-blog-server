// Package error
// created by lilei at 2021/11/23
package config

import "github.com/kataras/iris/v12"

// server
var (
	ServerIntervalError = ServerError{
		code:     10001,
		msg:      "服务器内部错误",
		httpCode: iris.StatusInternalServerError,
	}
	ServerDBError = ServerError{
		code:     10002,
		msg:      "数据库错误",
		httpCode: iris.StatusInternalServerError,
	}
	ServerCacheError = ServerError{
		code:     10003,
		msg:      "缓存错误",
		httpCode: iris.StatusInternalServerError,
	}
	ServerRPCConnectionError = ServerError{
		code:     10004,
		msg:      "请求服务连接错误",
		httpCode: iris.StatusInternalServerError,
	}
	ServerRPCBusinessError = ServerError{
		code:     10005,
		msg:      "请求服务业务错误",
		httpCode: iris.StatusInternalServerError,
	}
)

// client
var (
	ClientLoginRequiredError = ClientError{
		code:     20001,
		msg:      "用户操作需登录",
		httpCode: iris.StatusUnauthorized,
	}
	ClientParamValueInvalidError = ClientError{
		code:     20002,
		msg:      "参数值非法",
		httpCode: iris.StatusBadRequest,
	}
	ClientParamTypeInvalidError = ClientError{
		code:     20003,
		msg:      "参数类型非法",
		httpCode: iris.StatusBadRequest,
	}
	PermissionRequiredError = ClientError{
		code:     20004,
		msg:      "用户操作需授权",
		httpCode: iris.StatusForbidden,
	}
)

// business::user
var (
	UserNotExistsError = BusinessError{
		code:     30001,
		msg:      "用户不存在",
		httpCode: iris.StatusOK,
	}
	UserAlreadyExistsError = BusinessError{
		code:     30002,
		msg:      "用户已存在",
		httpCode: iris.StatusOK,
	}
)

// business::article
var (
	ArticleNotExistsError = BusinessError{
		code:     40001,
		msg:      "文章不存在",
		httpCode: iris.StatusOK,
	}
	ArticleAlreadyExistsError = BusinessError{
		code:     40002,
		msg:      "文章已存在",
		httpCode: iris.StatusOK,
	}
)
